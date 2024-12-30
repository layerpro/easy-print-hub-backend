# Build stage
FROM golang:1.23.3 AS builder

WORKDIR /app

# Copy go.mod and go.sum files first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go application
RUN go build -o main .

# Run stage
FROM debian:bookworm-slim

# Set a non-root user for better security
RUN useradd -m appuser
USER appuser

WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env

COPY wait-for-it.sh .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
