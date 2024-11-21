FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN apk add build-base

RUN go mod download

ENV CGO_ENABLED=1

RUN go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]