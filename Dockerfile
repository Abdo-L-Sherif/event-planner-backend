# =========================
# Stage 1: Build
# =========================
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Enable Go module proxy and sum database for faster downloads
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Copy dependency files first for better layer caching
COPY go.mod go.sum ./

# Download dependencies (cached if go.mod/go.sum unchanged)
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && go mod verify

# Copy source code
COPY . .

# Build static binary with optimizations
# -w: omit DWARF symbol table
# -s: omit symbol table and debug information
# -trimpath: remove file system paths from binary
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build \
    -ldflags="-w -s" \
    -trimpath \
    -o app .


# =========================
# Stage 2: Runtime (OpenShift-safe)
# =========================
FROM alpine:3.20

# TLS certs for HTTPS
RUN apk add --no-cache ca-certificates

# Writable directory for random UID
WORKDIR /app


# Copy binary
COPY --from=builder /app/app /app/app

# Copy .env file if it exists (optional)
COPY --from=builder /app/.env* /app/


# Ensure executable permissions
RUN chmod 755 /app/app

# Create a non-root user for OpenShift compatibility
# Note: OpenShift will override this UID, but it's good practice
RUN addgroup -g 1001 appgroup && \
    adduser -D -s /bin/sh -u 1001 -G appgroup appuser && \
    chown -R appuser:appgroup /app

USER appuser

# Expose port (OpenShift will use this for service configuration)
EXPOSE 8080

# Health check for OpenShift
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:${PORT:-8080}/health || exit 1


CMD ["/app/app"]
