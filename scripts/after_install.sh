#!/bin/bash
# Ensure the binary has correct permissions
chmod +x /usr/local/bin/awesomeProject

# Create necessary directories for logs
mkdir -p /var/log/awesomeProject
chown root:root /var/log/awesomeProject
chmod 755 /var/log/awesomeProject

# Copy .env file to the deployment directory (assuming it's included in artifacts)
if [ -f /tmp/.env ]; then
    cp /tmp/.env /usr/local/bin/.env
    chown root:root /usr/local/bin/.env
    chmod 600 /usr/local/bin/.env
fi

# Remove any stale lock files or temporary files
rm -f /tmp/awesomeProject*.lock