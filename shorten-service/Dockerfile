# -------- Stage 1: Build the Go binary --------
FROM golang:1.25 as builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Install git (required for some Go modules)
RUN apt-get update && apt-get install -y git

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o shorten-service ./cmd

# -------- Stage 2: Create lightweight runtime image --------
FROM alpine:latest

WORKDIR /root/

# Copy binary
COPY --from=builder /app/shorten-service .

# Copy environment file
COPY .env .

EXPOSE 8001

CMD ["./shorten-service"]
