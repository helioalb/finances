COMPOSE_FILE := deployments/compose.yml

.PHONY: up down down-purge logs ps restart

up:
	docker compose -f $(COMPOSE_FILE) up -d

down:
	docker compose -f $(COMPOSE_FILE) down

down-purge:
	docker compose -f $(COMPOSE_FILE) down --volumes --remove-orphans

logs:
	docker compose -f $(COMPOSE_FILE) logs -f

ps:
	docker compose -f $(COMPOSE_FILE) ps

restart: down up
