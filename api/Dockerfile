FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with optimization flags to reduce binary size by ~30-40%
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -trimpath \
    -o http-service ./main.go

# Final stage - minimal alpine image
FROM alpine:3.19

WORKDIR /app

# Install only essential runtime dependencies in one layer
RUN apk add --no-cache bash ca-certificates curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz \
    | tar xvz -C /usr/local/bin migrate && \
    chmod +x /usr/local/bin/migrate && \
    rm -rf /var/cache/apk/*

# Copy only the compiled binary and required files
COPY --from=builder /app/http-service .
COPY migration ./migration
COPY config ./config
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

# Run as non-root user for security
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser && \
    chown -R appuser:appuser /app

USER appuser

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]
