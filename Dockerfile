# Stage 1: Build the Go application
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o awesomeProject ./cmd/main.go

# Stage 2: Create the runtime image
FROM amazonlinux:2

# Install necessary dependencies
RUN yum update -y && yum install -y shadow-utils curl && yum clean all

# Create directories for the application and logs
RUN mkdir -p /usr/local/bin /var/log/awesomeProject
RUN chmod 755 /usr/local/bin /var/log/awesomeProject

# Copy the compiled binary from the builder stage
COPY --from=builder /app/awesomeProject /usr/local/bin/awesomeProject

# Copy the .env file and start_server.sh script
COPY .env /app/.env
COPY scripts/start_server.sh /usr/local/bin/start_server.sh

# Set permissions
RUN chmod +x /usr/local/bin/awesomeProject /usr/local/bin/start_server.sh
RUN chmod 600 .env /app/.env

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["/usr/local/bin/start_server.sh"]