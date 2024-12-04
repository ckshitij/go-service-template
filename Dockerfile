# Use the official Golang image as a builder stage
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the Go service with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -o go-template-service main.go

# Use a minimal base image for production
FROM alpine:latest

# Set the working directory inside the production container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/go-template-service .
COPY --from=builder /app/swagger ./swagger

# Expose the port the service listens on (default is 8080)
EXPOSE 8080

# Command to run the service
CMD ["./go-template-service"]
