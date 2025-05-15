#!/bin/bash
# Purpose: Validate that the application is running and accessible

# Wait for the application to start
sleep 10

# Check if the application process is running
if pgrep -f "awesomeProject" > /dev/null; then
    echo "Application is running"
    exit 0
else
    echo "Error: Application is not running"
    exit 1
fi