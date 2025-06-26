.PHONY: build run dev down clean logs test

# Build and run production containers
build: 
	docker-compose build

run: 
	docker-compose up -d

# Stop all containers
down:
    docker-compose down