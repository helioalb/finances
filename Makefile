COMPOSE_FILE := deployments/compose.yml

.PHONY: run build up down purge logs ps restart

run:
	docker compose -f $(COMPOSE_FILE) up -d

build:
	./scripts/build.sh

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
