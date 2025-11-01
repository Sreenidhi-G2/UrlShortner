# -------- Stage 1: Build the Go binary --------
FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o shorten-service ./cmd/main.go

# -------- Stage 2: Create a lightweight runtime image --------
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/shorten-service .
COPY .env .env
EXPOSE 8001
CMD ["./shorten-service"]
