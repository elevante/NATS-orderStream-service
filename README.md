<h1 align="center">NATS-orderStream-service</h1>

<p align="center">
  <img alt="Golang" src="https://img.shields.io/badge/Golang-74.6 %25-blue?style=for-the-badge&logo=appveyor">
  <img alt="HTML" src="https://img.shields.io/badge/HTML-14.3 %25-red?style=for-the-badge&logo=appveyor">
  <img alt="Makefile" src="https://img.shields.io/badge/Makefile-11.1 %25-green?style=for-the-badge&logo=appveyor">
</p>

## Description
NATS-orderStream-service is a Go service that processes order data. The service connects and subscribes to a channel in nats-streaming, writes the received data to a PostgreSQL database, and caches them in-memory. In case of service failure, the cache is restored from the database. The service also launches an HTTP server that outputs data by id from the cache.

## Installation and Launch
The project uses Docker and Makefile to simplify the installation and launch process. Here are the main commands:

```bash
# Building Docker images
make build

# Launching Docker containers
docker-compose up

# After launching Docker Compose, manually run 'model.sql' to create the necessary tables

# Running the Go service
make run

# Stopping Docker containers
docker-compose down


## Usage
After launching the service, you can get order data using its id. Just go to the following URL in your browser:

```
http://localhost:8080/orders/{id}
```
where `{id}` is the order id.
