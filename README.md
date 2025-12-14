# SES Go Boilerplate

A clean, scalable boilerplate for Go services with REST APIs, domain-driven structure, background workers, adapters, Git hooks, Swagger, and Makefile automation.

Project Structure
go-boilerplate/
│
├── cmd/
│ └── app/
│ ├── main.go
│ └── docs/ # Swagger documentation
│
├── internal/
│ ├── app/
│ │ ├── bootstrap/ # App initialization (DI, mongo, server)
│ │ ├── container.go # Dependency injection container
│ │ └── routes.go # Registers all domain routes
│ │
│ ├── health/ # Example domain module
│ │ ├── controller/
│ │ ├── model/
│ │ ├── processor/
│ │ ├── repository/
│ │ ├── routes/
│ │ └── service/
│ │
│ ├── workers/ # Background workers (MQTT, RabbitMQ, cron)
│ └── telemetry_worker.go
│
├── config/
│ ├── config.go # Unified config loader (JSON-based)
│ └── app.config.json # Primary config file
│
├── pkg/
│ ├── db/ # MongoDB connector
│ │ └── mongo.go
│ │
│ ├── redis/ # redis connector
│ │ ├── redis.go
│ │ └── helpers.go
│ │
│ ├── rabbitmq/ # rabbitmq connector
│ │ └── rabbitmq.go
│ │
│ ├── logger/ # Centralized logger
│ │ └── logger.go
│ │
│ └── utils/ # Shared helper utilities
│
├── tools/
│ └── git-hooks/
│ ├── install.sh
│ ├── pre-commit
│ └── pre-push
│
├── tests/ # Optional integration tests
│
├── .golangci.yml
├── .gitignore
├── Dockerfile
├── Makefile
└── README.md

## Features

- Organized domain-based architecture
- Clean separation: controller → service → repository
- Background workers for ingestion pipelines
- Adapters for MongoDB, Redis, RabbitMQ, MQTT
- Built-in Swagger generation
- Built-in Git hooks (linting, formatting, testing before push)
- Makefile for build/lint/test automation
- Easy to extend for new domains

## Requirements

- Go 1.21+
- MongoDB / Redis (optional)
- golangci-lint
- swag (Swagger generator)

## Run Application

```bash
go run ./cmd/app
make build
Run tests:
make test

# Generate Swagger
swag init -g cmd/app/main.go -o cmd/app/docs

# Git Hooks Setup
make hooks-install

# Run with docker
docker build -t go-boilerplate .
docker run -p 8080:8080 go-boilerplate
```

## API Documentation

```bash
swag init -g cmd/app/main.go -o cmd/app/docs
http://localhost:8080/swagger/index.html#/
```

## Adding a New Domain

Create a folder:

internal/<domain>/
controller/
model/
processor/
repository/
routes/
service/

Then register the domain in:

internal/app/container.go

internal/app/routes.go

## License

MIT
