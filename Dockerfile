# =========================
# Stage 1: Build
# =========================
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Only standard certs/git needed
RUN apk add --no-cache git ca-certificates tzdata

# DISABLE CGO (Pure Go for MySQL)
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# --- MEMORY OPTIMIZATIONS FOR OPENSHIFT BUILDER ---
# Forces garbage collection more often to save RAM
ENV GOGC=25
# Limits parallel compilation to 1 process to stop OOM kills
ENV GOMAXPROCS=1
# --------------------------------------------------

# Download dependencies separately for caching
COPY go.mod go.sum ./
RUN go mod download

# Build binary with minimal optimizations to save memory during build
COPY . .
RUN go build -ldflags="-s -w" -o app .

# =========================
# Stage 2: Runtime
# =========================
FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=builder /app/app /app/app

# Set permissions for OpenShift non-root policy
RUN chmod 755 /app/app && \
    addgroup -g 1001 appgroup && \
    adduser -D -s /bin/sh -u 1001 -G appgroup appuser && \
    chown -R appuser:appgroup /app

USER appuser
EXPOSE 8080

CMD ["/app/app"]
