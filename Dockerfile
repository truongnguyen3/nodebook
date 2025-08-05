# Multi-stage build for Nodebook
# Stage 1: Build frontend
FROM node:18-alpine AS frontend-builder

WORKDIR /app
COPY src/frontend/package*.json ./
RUN npm ci --only=production

COPY src/frontend/ ./
RUN npm run build

# Stage 2: Build Go application
FROM golang:1.24-alpine AS go-builder

# Install git (needed for some Go modules)
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Copy built frontend from previous stage
COPY --from=frontend-builder /app/../../dist/frontend ./dist/frontend

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o nodebook .

# Stage 3: Final runtime image
FROM alpine:latest

# Install Docker CLI (for docker mode execution)
RUN apk add --no-cache \
    ca-certificates \
    docker \
    tzdata \
    && rm -rf /var/cache/apk/*

WORKDIR /app

# Create user for security
RUN addgroup -g 1001 -S nodebook && \
    adduser -u 1001 -S nodebook -G nodebook

# Copy the binary from builder stage
COPY --from=go-builder /app/nodebook .

# Create notebooks directory
RUN mkdir -p /app/notebooks && \
    chown -R nodebook:nodebook /app

# Switch to non-root user
USER nodebook

# Expose port
EXPOSE 8000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8000/ || exit 1

# Default command - use Docker mode for code execution
CMD ["./nodebook", "--docker", "--bindaddress", "0.0.0.0", "/app/notebooks"]