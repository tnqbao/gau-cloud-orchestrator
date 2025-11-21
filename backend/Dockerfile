# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -a -installsuffix cgo \
    -o gau-cloud-service.bin \
    .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache \
    bash \
    ca-certificates \
    tzdata \
    curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz \
    | tar xvz -C /tmp \
    && mv /tmp/migrate /usr/local/bin/migrate \
    && chmod +x /usr/local/bin/migrate \
    && rm -rf /tmp/*

# Create non-root user for security
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Copy binary from builder
COPY --from=builder /app/gau-cloud-service.bin .

# Copy required files
COPY --chown=appuser:appuser migration ./migration
COPY --chown=appuser:appuser config ./config
COPY --chown=appuser:appuser entrypoint.sh .

# Make entrypoint executable
RUN chmod +x entrypoint.sh

# Switch to non-root user
USER appuser

# Expose port (can be overridden)
EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]