FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/main .
COPY docker/entrypoint.sh /root/entrypoint.sh

RUN chmod +x /root/entrypoint.sh

CMD ["/root/entrypoint.sh"]
