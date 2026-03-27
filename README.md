Go NLP Service

High-performance **Go microservice** for Natural Language Processing with gRPC integration to a Python ML backend.

Built as a production-ready example showcasing clean architecture, concurrency, observability, and real AI integration.

## Features

- ⚡ Fast REST API using **Gin**
- 🔄 High-concurrency processing with goroutines
- 📡 gRPC communication with Python NLP backend
- 🛡️ Rate limiting and graceful shutdown
- 📊 OpenTelemetry tracing support
- 🐳 Docker & Docker Compose ready
- 🧹 Clean architecture (`internal/` pattern)

## Quick Start

```bash
docker compose up --build
The service will be available at http://localhost:8080
Test the API
curl -X POST http://localhost:8080/api/v1/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "text": "I absolutely love this product, it changed my life!",
    "task": "sentiment"
  }'
Technologies
Go 1.23
Gin Framework
gRPC + Protobuf
OpenTelemetry
Docker & Docker Compose
Project Structure
cmd/api/           # Entry point
internal/
├── handler/       # HTTP handlers
├── service/       # Business logic
├── client/        # gRPC client to Python
├── model/         # Request/Response models
proto/             # Protocol Buffers
Made with ❤️ by Konstantin Lychkov
Looking for Senior Go Developer opportunities
Email: hotkeez@hotmail.com
WhatsApp: +1 (706) 998-7082
