#!/bin/bash

# Nodebook Deployment Script
# Usage: ./scripts/deploy.sh [docker|full|build-only]

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Functions
log() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
    exit 1
}

# Check if Docker is installed
check_docker() {
    if ! command -v docker &> /dev/null; then
        error "Docker is not installed. Please install Docker first."
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        error "Docker Compose is not installed. Please install Docker Compose first."
    fi
    
    if ! docker info &> /dev/null; then
        error "Docker daemon is not running. Please start Docker first."
    fi
}

# Build frontend
build_frontend() {
    log "Building frontend..."
    if [ -d "src/frontend" ]; then
        cd src/frontend
        if [ ! -d "node_modules" ]; then
            log "Installing frontend dependencies..."
            npm install
        fi
        log "Building frontend assets..."
        npm run build
        cd ../..
        success "Frontend built successfully"
    else
        warning "Frontend directory not found, skipping frontend build"
    fi
}

# Create notebooks directory
setup_notebooks() {
    log "Setting up notebooks directory..."
    mkdir -p notebooks
    
    # Copy example notebooks if they exist
    if [ -d "test/fixtures/notebooks" ]; then
        log "Copying example notebooks..."
        cp -r test/fixtures/notebooks/* notebooks/ 2>/dev/null || true
    fi
    
    success "Notebooks directory ready"
}

# Deploy with Docker mode
deploy_docker() {
    log "Deploying Nodebook in Docker mode..."
    
    check_docker
    build_frontend
    setup_notebooks
    
    log "Building Docker image..."
    docker-compose build
    
    log "Starting services..."
    docker-compose up -d
    
    success "Nodebook deployed successfully!"
    log "Access the application at: http://localhost:8000"
    log "View logs with: docker-compose logs -f"
}

# Deploy with full language support
deploy_full() {
    log "Deploying Nodebook with full language support..."
    
    check_docker
    build_frontend
    setup_notebooks
    
    # Build the Go binary first for full mode
    log "Building Go binary..."
    if ! go build -o nodebook .; then
        error "Failed to build Go binary"
    fi
    
    log "Building full Docker image (this may take a while)..."
    docker-compose -f docker-compose.full.yml build
    
    log "Starting services..."
    docker-compose -f docker-compose.full.yml up -d
    
    success "Nodebook (full) deployed successfully!"
    log "Access the application at: http://localhost:8000"
    log "View logs with: docker-compose -f docker-compose.full.yml logs -f"
}

# Build only (no deployment)
build_only() {
    log "Building Nodebook (no deployment)..."
    
    check_docker
    build_frontend
    
    log "Building Docker image..."
    docker-compose build
    
    success "Build completed successfully!"
    log "Deploy with: docker-compose up -d"
}

# Health check
health_check() {
    log "Checking application health..."
    
    # Wait for service to be ready
    local max_attempts=30
    local attempt=1
    
    while [ $attempt -le $max_attempts ]; do
        if curl -f http://localhost:8000/ &>/dev/null; then
            success "Application is healthy and responding"
            return 0
        fi
        
        log "Waiting for application to start... (attempt $attempt/$max_attempts)"
        sleep 2
        ((attempt++))
    done
    
    error "Application failed to start or is not responding"
}

# Show usage
usage() {
    echo "Usage: $0 [docker|full|build-only|health]"
    echo ""
    echo "Commands:"
    echo "  docker     Deploy with Docker mode (recommended, ~50MB)"
    echo "  full       Deploy with full language support (~2GB)"
    echo "  build-only Build Docker image without deploying"
    echo "  health     Check if deployed application is healthy"
    echo ""
    echo "Examples:"
    echo "  $0 docker    # Quick deployment with Docker execution"
    echo "  $0 full      # Full deployment with local language runtimes"
    echo "  $0 health    # Check application health"
}

# Main script
main() {
    case "${1:-docker}" in
        "docker")
            deploy_docker
            health_check
            ;;
        "full")
            deploy_full
            health_check
            ;;
        "build-only")
            build_only
            ;;
        "health")
            health_check
            ;;
        "--help"|"-h"|"help")
            usage
            ;;
        *)
            error "Unknown command: $1"
            usage
            ;;
    esac
}

# Run main function
main "$@"