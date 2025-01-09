FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

ENV CGO_ENABLED=1

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]