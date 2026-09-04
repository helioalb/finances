COMPOSE_FILE := deployments/compose.yml
GOCACHE ?= $(CURDIR)/.cache/go-build

.PHONY: run build test test-race up down purge logs ps restart

run:
	docker compose -f $(COMPOSE_FILE) up -d

build:
	./scripts/build.sh

test:
	GOCACHE="$(GOCACHE)" go test -race -v ./...

up:
	docker compose -f $(COMPOSE_FILE) up postgres -d

down:
	docker compose -f $(COMPOSE_FILE) down

purge:
	docker compose -f $(COMPOSE_FILE) down --volumes --remove-orphans

logs:
	docker compose -f $(COMPOSE_FILE) logs -f

ps:
	docker compose -f $(COMPOSE_FILE) ps

restart: down up

lint:
	docker run --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v2.13.2 golangci-lint run
