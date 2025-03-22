# Build Stage
FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

# Download dependencies and build
RUN go mod tidy && go build -o app

# Final Stage
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .

# Expose the port your app runs on
EXPOSE 8080
CMD ["./app exam"]
