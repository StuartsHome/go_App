# Docker

DOCKER_CMD = docker-compose
DOCKER_CMD_START = $(DOCKER_CMD) up -d server

.PHONY: docker-start
docker-start: docker-build
	$(DOCKER_CMD_START)

.PHONY: docker-build
docker-build:
	$(DOCKER_CMD) build

.PHONY: docker-stop
docker-stop:
	$(DOCKER_CMD) stop