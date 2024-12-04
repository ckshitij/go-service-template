# Default binary name and other variables
BINARY_NAME=go-template-service
DOCKER_IMAGE_NAME=go-template-service-image
DOCKER_TAG=0.0.1
DOCKERFILE=Dockerfile
DOCKER_COMPOSE_FILE=docker-compose.yml

.PHONY: help clean test build run docker-build docker-run docker-clean

# Default goal (help)
.DEFAULT_GOAL := help

# Git info (optional for versioning or labels)
GIT_COMMIT_HASH := $(shell git rev-list -1 HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

# Build-related commands
GOBUILD		=	go build
GOCLEAN		=	go clean
GOTEST		=	go test
INSTALL		=	go mod download

## Display help
help:
	@echo "Makefile commands:"
	@echo "  make build         - Build the Go binary"
	@echo "  make run           - Build and run the application locally"
	@echo "  make clean         - Clean the Go module cache and build artifacts"
	@echo "  make test          - Run tests"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"
	@echo "  make docker-clean  - Clean up Docker artifacts"

## Build the Go binary
build: ## Build the Go binary
	go mod tidy
	CGO_ENABLED=0 $(GOBUILD) -o ${BINARY_NAME}

## Run the Go binary locally
run: build ## Build and run locally
	./${BINARY_NAME}

## Clean Go build artifacts
clean: ## Clean Go build artifacts
	$(GOCLEAN)

## Run unit tests
test: ## Run unit tests
	$(GOTEST) -cover ./...

## Build Docker image
docker-build: ## Build Docker image
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) -f $(DOCKERFILE) .

## Run Docker container
docker-run: docker-build ## Build Docker image and run the container
	docker run --rm --network="host" -p 8080:8080 --env-file .env $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

## Clean up Docker images and containers
docker-clean: ## Remove Docker images and containers
	docker system prune -f
	docker rmi $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)
