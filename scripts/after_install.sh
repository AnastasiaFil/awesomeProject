#!/bin/bash
# Purpose: Perform post-installation tasks, such as setting permissions and verifying files

# Set correct permissions for the application binary and scripts
chmod +x /usr/local/bin/awesomeProject
chmod +x /usr/local/bin/validate_service.sh
chmod 600 /usr/local/bin/.env

# Verify the application binary exists
if [ ! -f /usr/local/bin/awesomeProject ]; then
    echo "Error: Application binary not found!"
    exit 1
fi

# Verify the validate_service.sh script exists
if [ ! -f /usr/local/bin/validate_service.sh ]; then
    echo "Error: start_server.sh script not found!"
    exit 1
fi

echo "AfterInstall completed successfully"
exit 0