#!/bin/sh

# Get service type from first argument, default to "http"
SERVICE_TYPE=${1:-http}

echo "Starting service: $SERVICE_TYPE"

# Run migrations first
echo "Running migrations..."
echo "Migration path: /app/migration"
echo "Database URL: $PGPOOL_URL"

# Use absolute path for migrations
migrate -database "$PGPOOL_URL" -path /app/migration up
if [ $? -ne 0 ]; then
    echo "Migrations failed. Exiting."
    exit 1
fi
echo "Migrations completed successfully."

# Start the HTTP API service
echo "Starting HTTP API service..."
if [ -f "./gau-cloud-service" ]; then
    ./gau-cloud-service
else
    echo "Binary gau-cloud-service not found. Exiting."
    exit 1
fi