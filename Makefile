# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker ps -a | grep -q template_db; then \
		echo "Container exists"; \
		if [ "$$(docker inspect -f '{{.State.Running}}' template_db)" = "false" ]; then \
			echo "Starting container..."; \
			docker start template_db; \
		fi \
	else \
		docker run --name template_db -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres; \
		echo "Container created and started"; \
	fi

# Shutdown DB container
docker-down:
	@if docker ps -a | grep -q template_db; then \
		echo "Stopping container..."; \
		docker stop template_db; \
		echo "Removing container..."; \
		docker rm template_db; \
	else \
		echo "Container does not exist"; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@go tool air --build.cmd "go build -o ./tmp/main cmd/api/main.go" --build.entrypoint "./tmp/main"

.PHONY: all build run test clean
