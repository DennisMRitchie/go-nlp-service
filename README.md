# Go NLP Service

! (https://img.shields.io/badge/Go-1.23-00ADD8.svg)
! (https://img.shields.io/badge/gRPC-00B4AB.svg)
! (https://img.shields.io/badge/Docker-2496ED.svg)

**High-performance Go microservice** for Natural Language Processing with gRPC integration to a Python ML backend.

Built as a production-ready example showcasing clean architecture, concurrency, observability, and real AI integration.

## Features

- ⚡ Fast REST API using **Gin**
- 🔄 High-concurrency with goroutines and channels
- 📡 gRPC communication with Python NLP backend
- 🛡️ Rate limiting and graceful shutdown
- 📊 OpenTelemetry tracing
- 🐳 Docker & Docker Compose ready
- 🧹 Clean architecture (`internal/` pattern)

## Quick Start

```bash
docker-compose up --build

API будет доступен на: http://localhost:8080

curl -X POST http://localhost:8080/api/v1/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "text": "I absolutely love this product, it changed my life!",
    "task": "sentiment"
  }'

Technologies
Go 1.23 • Gin • gRPC • Protobuf • OpenTelemetry • Docker
Project Structure

cmd/api/           # Entry point
internal/
├── handler/       # HTTP handlers
├── service/       # Business logic
├── client/        # gRPC client to Python
├── model/         # Request/Response models
proto/             # Protocol Buffers

Made with ❤️ by Konstantin Lychkov
Senior Go Developer | NLP Focus
Email: hotkeez@hotmail.com
WhatsApp: +1 (706) 998-7082
GitHub: (https://github.com/DennisMRitchie)
