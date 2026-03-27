# Go NLP Service

High-performance **Go microservice** for Natural Language Processing with gRPC integration to a Python ML backend.

Built as a production-ready example showcasing clean architecture, concurrency, observability, and real AI integration — perfect for demonstrating Go + NLP skills to recruiters and hiring managers.

## Features

- Fast REST API using Gin
- gRPC communication with Python NLP backend (sentiment analysis, text classification, entity recognition)
- Concurrent processing with goroutines
- Rate limiting
- OpenTelemetry tracing support
- Docker & Docker Compose ready
- Clean architecture

## Quick Start

### Using Docker Compose (recommended)

```bash
docker-compose up --build
