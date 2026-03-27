# Go NLP Service

! (https://img.shields.io/badge/Go-1.23-00ADD8.svg)
! (https://img.shields.io/badge/gRPC-00B4AB.svg)
! (https://img.shields.io/badge/Docker-2496ED.svg)

High-performance **Go microservice** for Natural Language Processing with gRPC integration to a Python ML backend.

Built as a production-ready example showcasing clean architecture, concurrency, observability, and real AI integration — perfect for demonstrating Go + NLP skills to recruiters and hiring managers.

## Features

- ⚡ Fast REST API using **Gin**
- 🔄 High-concurrency processing with goroutines
- 📡 gRPC communication with Python NLP backend
- 🛡️ Rate limiting & graceful shutdown
- 📊 OpenTelemetry tracing support
- 🐳 Docker & Docker Compose ready
- 🧹 Clean architecture (internal folder structure)

## Quick Start

```bash
# Clone and run
docker-compose up --build
