# Stage 1: Build the application
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wit .

# Stage 2: Create minimal runtime image
FROM alpine:latest

# Install necessary certificates
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the built binary from builder stage
COPY --from=builder /app/wit .

# Set the entrypoint
ENTRYPOINT ["./wit"]
