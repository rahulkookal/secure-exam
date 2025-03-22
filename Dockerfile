# Use the official Golang image as the build stage
FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

# Build the Go application
RUN go mod tidy && go build -o exam

# Use a minimal base image for final deployment
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/exam .

# Expose the application port
EXPOSE 8080

# Run the Go application
CMD ["./go-app"]
