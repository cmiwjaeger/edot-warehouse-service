# Use the official Golang image from the Docker Hub
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app/services/warehouse-service

# Copy the go.mod and go.sum files from the root directory
COPY go.mod go.sum ../../
COPY shared/ ../../shared

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code of the service into the container
COPY ./services/warehouse-service .

# Build the Go app
RUN GOARCH=amd64 GOOS=linux go build -o main ./cmd/web/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/services/warehouse-service/main /app

# Expose port 3100 to the outside world
EXPOSE 3100

# Command to run the executable
CMD ["./main"]

