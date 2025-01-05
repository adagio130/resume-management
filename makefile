APP_NAME =  resume-app
DOCKER_COMPOSE_FILE = docker-compose.yml

# Default target
.DEFAULT_GOAL := help

# Build the Docker Compose services
build:
	@echo "Building Docker Compose services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

# Run the application
run:
	@echo "Starting the application using Docker Compose..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop the running application
stop:
	@echo "Stopping the application..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Clean up Docker resources
clean:
	@echo "Stopping and cleaning up resources..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --rmi all --volumes --remove-orphans

# View logs of the running services
logs:
	@echo "Fetching logs from the application..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

# Restart the application (stop + run)
restart: stop run

# Show the help menu
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build      Build the Docker Compose services."
	@echo "  run        Start the application using Docker Compose."
	@echo "  stop       Stop the running application."
	@echo "  clean      Stop and remove containers, images, volumes, and networks."
	@echo "  logs       View logs of the running services."
	@echo "  restart    Restart the application."
	@echo "  help       Show this help message."