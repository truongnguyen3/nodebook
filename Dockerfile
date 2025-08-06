# Multi-stage build for Nodebook
# Stage 1: Build Go application
FROM golang:1.24-alpine AS go-builder

# Install git (needed for some Go modules)
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and pre-built frontend
COPY . .

# Verify frontend is already built
RUN ls -la dist/frontend/ || (echo "Error: Frontend not found in dist/frontend/. Please build frontend first." && exit 1)

# Build the Go binary for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o nodebook .

# Stage 2: Runtime image with language support
FROM alpine:latest

# Install necessary packages including language runtimes
RUN apk add --no-cache \
    ca-certificates \
    tzdata \
    python3 \
    py3-pip \
    nodejs \
    npm \
    openjdk17-jre \
    php \
    ruby \
    lua5.3 \
    gcc \
    g++ \
    rust \
    curl \
    wget \
    && rm -rf /var/cache/apk/*

WORKDIR /app

# Create user for security
RUN addgroup -g 1001 -S nodebook && \
    adduser -u 1001 -S nodebook -G nodebook

# Copy the Linux binary from builder stage
COPY --from=go-builder /app/nodebook .
COPY --from=go-builder /app/dist/frontend ./dist/frontend
COPY --from=go-builder /app/src/recipes ./src/recipes

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

# Default command - web mode (local execution since we're in a container)
CMD ["./nodebook", "web", "--bindaddress", "0.0.0.0", "--port", "8000", "/app/notebooks"]