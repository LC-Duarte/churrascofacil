# Use the official Golang image as a base
FROM golang:1.17-alpine AS builder
# Set the working directory
WORKDIR /churrascofacil
# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy the rest of the application source code
COPY . .
# Build the application
RUN go build -o bin/v0.1/cfserver cmd/cfserver/main.go
# Use a minimal Alpine-based image to reduce the final image size
FROM alpine:3.14
# Set the working directory
WORKDIR /churrascofacil
# Copy the server binary from the builder image
RUN mkdir data
RUN mkdir bin
RUN mkdir bin/v0.1
COPY --from=builder /churrascofacil/bin/v0.1/cfserver bin/v0.1/

# Start the server
WORKDIR /churrascofacil/bin/v0.1/
# Expose the server port
EXPOSE 8080
CMD ["./churrascofacil"]