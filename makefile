# Makefile for Docker Compose

# Define the Docker Compose project name
PROJECT_NAME := api-simulator

# Docker Compose commands
COMPOSE := docker-compose -p $(PROJECT_NAME)

# Get the Go version from the system
GO_VERSION := $(shell go version | awk -F'[ .]' '{print $$3"."$$4}' | sed 's/^go//')

# Targets

# Build the Docker Compose project
create:
	@echo "Building the Docker Compose project"
	$(COMPOSE) build

# Run the Docker Compose project
up:
	@echo "Running the Docker Compose project in detached mode"
	$(COMPOSE) up -d

# Stop the Docker Compose project
down:
	@echo "Stopping the Docker Compose project"
	$(COMPOSE) down

# Remove the Docker Compose project (including volumes)
remove:
	@echo "Removing the Docker Compose project (including volumes)"
	$(COMPOSE) down -v

# Run the Docker Compose project, build it if needed
run: create up

# Remove and then build and run the Docker Compose project
rebuild: remove run

.PHONY: set-go-version tidy

# Set Go version in go.mod
set-go-version:
	@echo "Setting Go version to $(GO_VERSION) in go.mod"
	@go mod edit -go=$(GO_VERSION)

# Run 'go mod tidy' to ensure dependencies match the Go version
tidy:
	@echo "Running 'go mod tidy'"
	@go mod tidy

# Set Go version and run 'go mod tidy' in one step
update-go: set-go-version tidy

# Help target to display available commands
help:
	@echo "Available targets:"
	@echo "  create           Build the Docker Compose project"
	@echo "  up               Run the Docker Compose project in detached mode"
	@echo "  down             Stop the Docker Compose project"
	@echo "  remove           Remove the Docker Compose project (including volumes)"
	@echo "  run              Remove, build, and run the Docker Compose project"
	@echo "  rebuild          Remove, build, and run the Docker Compose project (rebuild from scratch)"
	@echo "  set-go-version   Set the Go version in go.mod"
	@echo "  tidy             Run 'go mod tidy'"
	@echo "  update-go        Set Go version and run 'go mod tidy'"
	@echo "  help             Display this help message (current target)"
