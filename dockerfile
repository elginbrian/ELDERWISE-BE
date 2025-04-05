FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -o elderwise cmd/main.go && \
    go build -o network_check scripts/network_check.go && \
    chmod +x scripts/entrypoint.sh

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata wget

WORKDIR /app
VOLUME ["/app/logs"]

COPY --from=builder /app/elderwise .
COPY --from=builder /app/network_check .
COPY --from=builder /app/scripts/entrypoint.sh .

RUN chmod +x /app/entrypoint.sh /app/elderwise /app/network_check

EXPOSE 3000

ENTRYPOINT ["/app/entrypoint.sh"]
