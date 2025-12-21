# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Run the application (Lightweight image)
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]