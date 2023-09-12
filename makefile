# Makefile to build and run a Docker container from a specified Dockerfile

# Define the name for your Docker image
IMAGE_NAME = "simulation:latest"

# Define the port number to expose in the container and on the host
PORT = 9999

# Define the Docker run command with the necessary options
DOCKER_RUN = docker run -d -p $(PORT):$(PORT) --rm $(IMAGE_NAME)

# Targets

# Build the Docker image from the specified Dockerfile
create:
	docker build -t $(IMAGE_NAME) .

# Run the Docker container
run: build
	$(DOCKER_RUN)
# Stop the Docker container
stop:
	docker stop $$(docker ps -q --filter ancestor=$(IMAGE_NAME))

# Remove the Docker image
remove:
	docker rmi $(IMAGE_NAME)

# Clean up Docker containers and images
clean:
	docker stop $$(docker ps -a -q)  # Stop all containers
	docker rm $$(docker ps -a -q)    # Remove all containers
	docker rmi $$(docker images -q)  # Remove all images
