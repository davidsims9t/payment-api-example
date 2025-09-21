FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o payments-service main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/payments-service .
EXPOSE 4002
CMD ["./payments-service"]