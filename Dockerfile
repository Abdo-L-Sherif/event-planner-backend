# =========================
# Stage 1: Build
# =========================
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Only standard certs/git needed; no GCC required for pure Go MySQL
RUN apk add --no-cache git ca-certificates tzdata

# Disable CGO to ensure a static, pure-Go binary
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build binary
COPY . .
RUN go build -ldflags="-s -w" -o app .

# =========================
# Stage 2: Runtime
# =========================
FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=builder /app/app /app/app

# OpenShift Security: Set permissions and non-root user
RUN chmod 755 /app/app && \
    addgroup -g 1001 appgroup && \
    adduser -D -s /bin/sh -u 1001 -G appgroup appuser && \
    chown -R appuser:appgroup /app

USER appuser
EXPOSE 8080

CMD ["/app/app"]
