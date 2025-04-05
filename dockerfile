FROM golang:1.23.4 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -o elderwise cmd/main.go && \
    chmod +x scripts/entrypoint.sh

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
VOLUME ["/app"]

COPY --from=builder /app/elderwise .
COPY --from=builder /app/scripts/entrypoint.sh .
COPY --from=builder /app/scripts/network_check.go ./scripts/

RUN chmod +x /app/entrypoint.sh /app/elderwise

EXPOSE 3000

ENTRYPOINT ["/app/entrypoint.sh"]
