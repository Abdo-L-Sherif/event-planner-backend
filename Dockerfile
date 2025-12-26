FROM golang:1.24-alpine AS builder
WORKDIR /app
# Install build-base for gcc and musl-dev for C headers
RUN apk add --no-cache git ca-certificates tzdata build-base musl-dev

ENV GOPROXY=https://proxy.golang.org,direct
# CHANGE: Set CGO_ENABLED to 1
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags="-s" -o app .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/app /app/app
RUN chmod 755 /app/app
# Create non-root user
RUN addgroup -g 1001 appgroup && \
    adduser -D -s /bin/sh -u 1001 -G appgroup appuser && \
    chown -R appuser:appgroup /app
USER appuser
EXPOSE 8080
CMD ["/app/app"]
