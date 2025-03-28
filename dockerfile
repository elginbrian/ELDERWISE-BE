FROM golang:1.23.4 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o elderwise cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/elderwise .

EXPOSE 3000

CMD ["./elderwise"]
