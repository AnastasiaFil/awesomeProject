#!/bin/bash
# Purpose: Start the awesomeProject application

# Load environment variables from .env file
if [ -f /usr/local/bin/.env ]; then
    export $(grep -v '^#' /usr/local/bin/.env | xargs)
else
    echo "Error: .env file not found!"
    exit 1
fi

# Start the Go application
echo "Starting awesomeProject application..."
/usr/local/bin/awesomeProject &

# Capture the process ID
APP_PID=$!
echo "Application started with PID $APP_PID"

# Wait briefly to ensure the application starts
sleep 5

# Check if the application is still running
if kill -0 $APP_PID 2>/dev/null; then
    echo "Application is running"
    exit 0
else
    echo "Error: Application failed to start"
    exit 1
fi