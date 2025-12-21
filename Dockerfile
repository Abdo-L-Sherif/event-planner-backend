# =========================
# Stage 1: Build
# =========================
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app


# =========================
# Stage 2: Runtime
# =========================
FROM alpine:3.20

# Install CA certs for HTTPS
RUN apk add --no-cache ca-certificates

# Create app directory that random UID can access
WORKDIR /app

# Copy binary
COPY --from=builder /app/app /app/app

# Make sure it's executable
RUN chmod 755 /app/app

# OpenShift-friendly port
EXPOSE 8080

# Run as non-root (OpenShift will override UID anyway)
USER 1001

CMD ["/app/app"]
