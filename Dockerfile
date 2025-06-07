FROM golang:1.24 AS builder

WORKDIR /app
COPY . .
RUN go build -o irrig8r ./cmd/main.go

FROM debian:bookworm-slim
COPY --from=builder /app/irrig8r /usr/local/bin/irrig8r
ENTRYPOINT ["/usr/local/bin/irrig8r"]
