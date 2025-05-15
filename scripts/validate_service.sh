#!/bin/bash
# Purpose: Validate that the application is running and accessible

# Wait for the application to start (adjust sleep time as needed)
sleep 10

# Check if the application is running by querying the health check endpoint
# Replace with your actual health check endpoint (e.g., /health)
HEALTH_CHECK_URL="http://localhost:8080/health"
RESPONSE=$(curl --silent --write-out "%{http_code}" --output /dev/null $HEALTH_CHECK_URL)

if [ "$RESPONSE" -eq 200 ]; then
    echo "Application is running and health check passed"
    exit 0
else
    echo "Error: Health check failed with status code $RESPONSE"
    exit 1
fi