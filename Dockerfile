# =========================
# Stage 1: Build
# =========================
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app


# =========================
# Stage 2: Runtime (OpenShift-safe)
# =========================
FROM alpine:3.20

# TLS certs for HTTPS
RUN apk add --no-cache ca-certificates

# Writable directory for random UID
WORKDIR /app

# Copy binary and env file
COPY --from=builder /app/app /app/app
COPY --from=builder /app/.env /app/.env

# Ensure executable permissions
RUN chmod 755 /app/app

# OpenShift-friendly port
EXPOSE 8080

# Non-root user (OpenShift will override UID anyway)
USER 1001

CMD ["/app/app"]
