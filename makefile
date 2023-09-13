# Makefile for Docker Compose

# Define the Docker Compose project name
PROJECT_NAME := api-simulator

# Docker Compose commands
COMPOSE := docker-compose -p $(PROJECT_NAME)

# Targets

# Build the Docker Compose project
create:
	$(COMPOSE) build

# Run the Docker Compose project
up:
	$(COMPOSE) up -d

# Stop the Docker Compose project
down:
	$(COMPOSE) down

# Remove the Docker Compose project (including volumes)
remove:
	$(COMPOSE) down -v

# Run the Docker Compose project, build it if needed
run: create up

# Remove and then build and run the Docker Compose project
rebuild: remove run

# By default, run the Docker Compose project
.DEFAULT_GOAL := run
