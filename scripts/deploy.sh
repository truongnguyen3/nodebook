#!/bin/bash

# Nodebook Deployment Script
# Usage: ./scripts/deploy.sh [docker|full|build-only]
# Note: Frontend should be pre-built in dist/frontend/ directory

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

# Check frontend build
check_frontend() {
    log "Checking frontend build..."
    if [ -d "dist/frontend" ]; then
        success "Frontend already built and available in dist/frontend"
    else
        warning "Frontend build not found in dist/frontend"
        if [ -d "src/frontend" ]; then
            log "Frontend source found, but build is missing"
            log "Please run 'make build-frontend' to build the frontend first"
            error "Frontend build required but not found"
        else
            error "Neither frontend build nor source found"
        fi
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
    check_frontend
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
    check_frontend
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
    check_frontend
    
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
    echo "Usage: $0 [docker|full|build-frontend|build-only|health]"
    echo ""
    echo "Commands:"
    echo "  docker         Deploy with Docker mode (recommended, ~50MB)"
    echo "  full           Deploy with full language support (~2GB)"
    echo "  build-frontend Deploy with frontend building (Node 16, ~80MB)"
    echo "  build-only     Build Docker image without deploying"
    echo "  health         Check if deployed application is healthy"
    echo ""
    echo "Prerequisites:"
    echo "  - Frontend pre-built in dist/frontend/ (for docker/full modes)"
    echo "  - Docker and Docker Compose must be installed"
    echo ""
    echo "Examples:"
    echo "  $0 docker          # Quick deployment with pre-built frontend"
    echo "  $0 build-frontend  # Build frontend from source with Node 16"
    echo "  $0 full            # Full deployment with local language runtimes"
    echo "  $0 health          # Check application health"
}

# Deploy with frontend building (Node 16)
deploy_build_frontend() {
    log "Deploying Nodebook with frontend building (Node 16)..."
    
    check_docker
    setup_notebooks
    
    log "Building Docker image with frontend build..."
    docker-compose -f docker-compose.with-frontend-build.yml build
    
    log "Starting services..."
    docker-compose -f docker-compose.with-frontend-build.yml up -d
    
    success "Nodebook (with frontend build) deployed successfully!"
    log "Access the application at: http://localhost:8000"
    log "View logs with: docker-compose -f docker-compose.with-frontend-build.yml logs -f"
    
    health_check
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
        "build-frontend")
            deploy_build_frontend
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