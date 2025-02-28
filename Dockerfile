# Build stage - updated to Go 1.22
FROM golang:1.22-alpine AS builder

# Install git and necessary build tools
RUN apk add --no-cache git ca-certificates tzdata alpine-sdk

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go/bin/app .

# Final stage
FROM alpine:3.18

# Add necessary certificates and timezone data
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Create a non-root user to run the application
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /go/bin/app .

# Switch to non-root user
USER appuser

# Expose port the application will run on
EXPOSE 8080

# Run the application with the server command
ENTRYPOINT ["/app/app", "exam"]