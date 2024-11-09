# Pasipiya Go boilerplate

This is a boilerplate Go project for RESTful API applications, complete with logging, configuration management, and a Dockerfile for containerization.

## Structure

- `cmd/` contains the main application entry point.
- `config/` handles app configuration.
- `internal/` contains internal packages, such as server setup, handlers, services, and repositories.
- `pkg/` holds reusable packages like `logger`.
- `scripts/` includes setup and deployment scripts.
- `tests/` for integration and unit tests.

## Getting Started

1. Clone the repository.
2. Install dependencies:

   ```bash
   go mod download

3. Run Applicartion

    ```bash
    go run cmd/app/main.go

4. Run with docker

   ```bash
    docker build -t pasipiya-go-boilerplate .
    docker run -p 8080:8080 pasipiya-go-boilerplate
