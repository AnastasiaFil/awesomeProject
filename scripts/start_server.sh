#!/bin/bash
# Load environment variables from .env file
if [ -f /usr/local/bin/.env ]; then
    export $(grep -v '^#' /usr/local/bin/.env | xargs)
fi

# Start the Go application
nohup /usr/local/bin/awesomeProject > /var/log/awesomeProject/app.log 2>&1 &

# Verify the application is running
sleep 5
if pgrep -f awesomeProject > /dev/null; then
    echo "awesomeProject started successfully"
else
    echo "Failed to start awesomeProject"
    cat /var/log/awesomeProject/app.log
    exit 1
fi