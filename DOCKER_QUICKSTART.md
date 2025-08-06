# 🐳 Nodebook Docker Quick Start

Get Nodebook running in Docker in under 2 minutes!

## ⚡ Super Quick Start

```bash
# 1. Clone and enter directory
git clone <repository> && cd nodebook

# 2. Frontend is already built in dist/frontend/ ✅

# 3. Deploy with one command
./scripts/deploy.sh docker

# 4. Open in browser
open http://localhost:8000
```

That's it! 🎉

## 📋 What You Get

✅ **Multi-language REPL** - Support for 20+ programming languages  
✅ **Web Interface** - Clean, modern React UI  
✅ **Docker Execution** - Isolated, secure code execution  
✅ **Auto-scaling** - Containers created/destroyed per execution  
✅ **Persistent Storage** - Notebooks saved to `./notebooks/`  
✅ **Production Ready** - Health checks, logging, monitoring  

## 🎯 Available Commands

```bash
# Standard deployment (recommended)
./scripts/deploy.sh docker

# Full language support (larger image)
./scripts/deploy.sh full

# Build only (no deployment)
./scripts/deploy.sh build-only

# Check health
./scripts/deploy.sh health
```

## 🛠️ Manual Commands

```bash
# Build and run
docker-compose up -d

# View logs
docker-compose logs -f

# Stop
docker-compose down

# Full rebuild
docker-compose build --no-cache
```

## 📁 File Structure

```
nodebook/
├── Dockerfile              # Main Docker image (Docker mode)
├── Dockerfile.full         # Full image with language runtimes
├── docker-compose.yml      # Standard deployment
├── docker-compose.full.yml # Full language deployment
├── .dockerignore           # Docker build optimization
├── nginx.conf              # Production proxy config
├── scripts/deploy.sh       # One-click deployment
├── DEPLOYMENT.md           # Detailed deployment guide
└── notebooks/              # Your notebook files (created)
```

## 🔧 Configuration

### Prerequisites
- ✅ **Frontend pre-built** in `dist/frontend/` directory
- ✅ **Docker & Docker Compose** installed
- ✅ **Go 1.24** binary available

### Ports
- **8000** - Main application port
- **80** - Nginx proxy (production)

### Volumes
- `./notebooks` → `/app/notebooks` - Notebook storage
- `/var/run/docker.sock` → `/var/run/docker.sock` - Docker access

### Environment
- Docker mode enabled by default
- Binds to `0.0.0.0:8000` in container
- Non-root user for security

## 🌐 Language Support

**Fully Supported Languages:**
- Python 3 🐍
- Node.js/JavaScript 🟨
- Go 🐹
- Java ☕
- TypeScript 🔷
- Swift 🦉
- Rust 🦀
- Ruby 💎
- PHP 🐘
- C/C++ ⚙️
- And 10+ more!

## 🚀 Production Tips

1. **Use Docker mode** (default) for security
2. **Mount persistent volumes** for notebooks
3. **Set up reverse proxy** with nginx.conf
4. **Enable SSL/TLS** for HTTPS
5. **Monitor resource usage**
6. **Regular backups** of notebooks directory

## 🆘 Quick Troubleshooting

| Problem | Solution |
|---------|----------|
| Port 8000 in use | Change port in docker-compose.yml |
| Docker permission denied | Add user to docker group |
| Build fails | Run `docker system prune` and retry |
| Can't access externally | Change bind address to `0.0.0.0` |

## 📞 Getting Help

- 📖 [Full Deployment Guide](DEPLOYMENT.md)
- 📋 [Main README](README.md)
- 🐛 [Issue Tracker](../../issues)

---

**Happy Coding!** 🚀