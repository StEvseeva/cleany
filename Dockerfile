# Build stage
FROM golang:1.24-alpine AS builder

# Install git, ca-certificates, and build tools (needed for go mod download and CGO)
RUN apk add --no-cache git ca-certificates build-base && \
    apk upgrade --no-cache

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage
FROM alpine:3.19

# Install ca-certificates for HTTPS requests and upgrade packages
RUN apk --no-cache add ca-certificates && \
    apk --no-cache add postgresql-client && \
    apk upgrade --no-cache && \
    rm -rf /var/cache/apk/*

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy any additional files if needed (like config files)
COPY --from=builder /app/docs ./docs

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Set read-only root filesystem
# Note: This requires the app to write to /tmp or other writable directories

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/rooms || exit 1

# Run the application
CMD ["./main"] 