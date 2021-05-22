# Introduction

This service is makes it possible to store problem reports via an API using graphql

# Building and tagging with Docker

`docker build -f deployments/Dockerfile -t diwise/api-problemreport:latest .`

# Build for local testing with Docker Compose

`docker compose -f ./deployments/docker-compose.yml build`

# Running locally with Docker Compose

`docker compose -f ./deployments/docker-compose.yml up`

The ingress service will exit fatally and restart a couple of times until the RabbitMQ container is properly initialized and ready to accept connections. This is to be expected.

# Clean up the environment

`docker compose -f ./deployments/docker-compose.yml down -v`

To clean up the environment properly after testing.

# Regenerate GraphQL files

`go run github.com/99designs/gqlgen -v -c internal/pkg/graphql/gqlgen.yml`
