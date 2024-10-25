#light weight go image
FROM golang:1.23-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=1

RUN apk add --no-cache gcc libc-dev

COPY go.mod go.sum ./


RUN go mod download

COPY . .

RUN go build -o simple-go-crud main.go


FROM alpine:latest

RUN apk add --no-cache sqlite 

WORKDIR /app

COPY --from=builder /app/simple-go-crud .

EXPOSE 8080

CMD ["./simple-go-crud"]