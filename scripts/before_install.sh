#!/bin/bash
# Purpose: Prepare the environment before installing the new application version

# Ensure the log directory exists
mkdir -p /var/log/awesomeProject
chmod 755 /var/log/awesomeProject

# Clean up any previous application artifacts if necessary
if [ -f /usr/local/bin/awesomeProject ]; then
    echo "Removing old application binary..."
    rm -f /usr/local/bin/awesomeProject
fi

# Ensure the .env file is in place (optional, as it's copied in Dockerfile)
if [ ! -f /usr/local/bin/.env ]; then
    echo "Error: .env file not found!"
    exit 1
fi

echo "BeforeInstall completed successfully"
exit 0