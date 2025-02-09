# Start from the official Golang image
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
#RUN go build -o server ./cmd/blog-backend
RUN CGO_ENABLED=0 go build -o /out/bin/server ./cmd/*

# Start with a smaller base image for the final image
FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /out/bin/server /app/server

# Expose the port your server runs on
EXPOSE 8080

# Run the server
CMD ["./server"]
