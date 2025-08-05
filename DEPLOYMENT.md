# 🐳 Nodebook Docker Deployment Guide

This guide explains how to deploy Nodebook using Docker in various configurations.

## 🚀 Quick Start

### Option 1: Docker Mode (Recommended) - Lightweight ~50MB

```bash
# 1. Build and run with docker-compose
docker-compose up -d

# 2. Access the application
open http://localhost:8000
```

### Option 2: Full Language Support - Large ~2GB

```bash
# 1. Build and run with full language support
docker-compose -f docker-compose.full.yml up -d

# 2. Access the application  
open http://localhost:8000
```

## 🏗️ Build Options

### Standard Build (Docker Mode)
- **Size**: ~50MB
- **Execution**: Code runs in separate Docker containers
- **Languages**: All 20+ languages supported via Docker
- **Performance**: Fast, isolated execution
- **Security**: Maximum isolation

```bash
docker build -t nodebook:latest .
```

### Full Build (Local Runtimes)
- **Size**: ~2GB
- **Execution**: Code runs directly in the main container
- **Languages**: Pre-installed runtimes
- **Performance**: Faster startup, no container overhead
- **Security**: Less isolation

```bash
docker build -f Dockerfile.full -t nodebook:full .
```

## 🔧 Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8000` | HTTP server port |
| `BIND_ADDRESS` | `0.0.0.0` | Bind address |
| `NOTEBOOKS_PATH` | `/app/notebooks` | Notebooks directory |
| `DOCKER_MODE` | `true` | Enable Docker execution |

### Volume Mounts

| Host Path | Container Path | Purpose |
|-----------|----------------|---------|
| `./notebooks` | `/app/notebooks` | User notebooks storage |
| `/var/run/docker.sock` | `/var/run/docker.sock` | Docker socket (Docker mode only) |

## 🚢 Production Deployment

### Using Docker Compose

```yaml
version: '3.8'
services:
  nodebook:
    image: nodebook:latest
    ports:
      - "8000:8000"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - ./notebooks:/app/notebooks
    environment:
      - DOCKER_HOST=unix:///var/run/docker.sock
    restart: unless-stopped
```

### Using Kubernetes

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nodebook
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nodebook
  template:
    metadata:
      labels:
        app: nodebook
    spec:
      containers:
      - name: nodebook
        image: nodebook:latest
        ports:
        - containerPort: 8000
        volumeMounts:
        - name: notebooks
          mountPath: /app/notebooks
        - name: docker-sock
          mountPath: /var/run/docker.sock
      volumes:
      - name: notebooks
        persistentVolumeClaim:
          claimName: nodebook-notebooks
      - name: docker-sock
        hostPath:
          path: /var/run/docker.sock
---
apiVersion: v1
kind: Service
metadata:
  name: nodebook-service
spec:
  selector:
    app: nodebook
  ports:
  - port: 80
    targetPort: 8000
  type: LoadBalancer
```

## 🔒 Security Considerations

### Docker Mode (Recommended)
- ✅ Code execution is isolated in separate containers
- ✅ Main container runs as non-root user
- ✅ Minimal attack surface
- ✅ Automatic cleanup of execution containers

### Local Mode
- ⚠️ Code runs directly in main container
- ⚠️ Requires more security hardening
- ⚠️ Potential for resource exhaustion
- ✅ Still runs as non-root user

### Production Recommendations
1. **Use Docker mode** for code execution
2. **Run behind a reverse proxy** (nginx, traefik)
3. **Set up SSL/TLS** certificates
4. **Limit resource usage** with Docker constraints
5. **Regular security updates** for base images
6. **Network isolation** with Docker networks
7. **Backup notebooks** directory regularly

## 🛠️ Maintenance

### Updating the Application

```bash
# Pull latest changes
git pull origin main

# Rebuild and restart
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```

### Cleaning Up

```bash
# Remove unused Docker images
docker system prune -a

# Clean up execution containers (Docker mode)
docker container prune
```

### Logs and Monitoring

```bash
# View logs
docker-compose logs -f nodebook

# Monitor resource usage
docker stats

# Health check
curl http://localhost:8000/
```

## 🆘 Troubleshooting

### Common Issues

1. **Port already in use**
   ```bash
   # Change port in docker-compose.yml
   ports:
     - "8080:8000"  # Use port 8080 instead
   ```

2. **Docker socket permission denied**
   ```bash
   # Add user to docker group
   sudo usermod -aG docker $USER
   # Restart session
   ```

3. **Notebooks not persisting**
   ```bash
   # Check volume mount
   docker-compose exec nodebook ls -la /app/notebooks
   ```

4. **Frontend not loading**
   ```bash
   # Rebuild with no cache
   docker-compose build --no-cache
   ```

### Performance Tuning

```yaml
# Add to docker-compose.yml service
deploy:
  resources:
    limits:
      cpus: '2.0'
      memory: 1G
    reservations:
      cpus: '0.5'
      memory: 256M
```

## 📚 Additional Resources

- [Nodebook Documentation](../README.md)
- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
- [Docker Compose Reference](https://docs.docker.com/compose/compose-file/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)