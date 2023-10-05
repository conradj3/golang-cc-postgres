.PHONY: all build up down

# Build all services
build: docker-compose build

# Spin up all services
up: docker-compose up

# Spin down all services
down: docker-compose down

# Build and spin up the producer
up-producer: docker-compose up producer

# Build and spin up the consumer
up-consumer: docker-compose up consumer
