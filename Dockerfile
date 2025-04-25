# Start from the official Golang base image
FROM golang:1.23-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    DB_HOST=109.106.253.187 \
    DB_PORT=3306 \
    DB_USER=u1275606_user_backend_ld_project \
    DB_PASSWORD=LD_db1234! \
    DB_NAME=u1275606_backend_ld_project \
    SDK_LD=sdk-d4a84fa2-23f0-486b-af59-8985243ede13

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main .

# Final stage: minimal image
FROM alpine:latest

# Set working directory in the container
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/main .

# Expose application port
EXPOSE 3000

# Run the binary
CMD ["./main"]
