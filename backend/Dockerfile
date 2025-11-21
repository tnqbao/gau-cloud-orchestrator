### Stage 1: Builder
FROM --platform=linux/amd64 golang:1.25 AS builder

WORKDIR /app

# Faster Go fetch
ENV GOPROXY=https://proxy.golang.org,direct

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the HTTP service
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o gau-cloud-service.bin ./http

# Build consumer
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o gau-cloud-consumer.bin ./consumer

### Stage 2: Migrate binary
FROM alpine:3.18 AS migrator

RUN apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz \
       -o /tmp/migrate.tar.gz \
    && tar -xzf /tmp/migrate.tar.gz -C /tmp \
    && chmod +x /tmp/migrate \
    && rm -rf /var/cache/apk/* /tmp/*.tar.gz


### Stage 3: Runtime
FROM alpine:3.18

WORKDIR /app

# Runtime dependencies
RUN apk add --no-cache bash ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Copy binaries
COPY --from=builder /app/gau-cloud-service.bin .
# COPY --from=builder /app/gau-cloud-consumer.bin .  # nếu cần
COPY --from=migrator /tmp/migrate /usr/local/bin/migrate

# Copy assets
COPY --chown=appuser:appuser migration ./migration
COPY --chown=appuser:appuser config ./config
COPY --chown=appuser:appuser entrypoint.sh .
RUN chmod +x entrypoint.sh

# Switch to non-root
USER appuser

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]
