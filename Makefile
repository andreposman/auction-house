# Auction House API - Simple Makefile

.PHONY: help db-start db-stop run dev migrate

help: ## Show available commands
	@echo "Available commands:"
	@echo "  make db-start - Start database"
	@echo "  make db-stop  - Stop database"
	@echo "  make run      - Run the app with air (hot reload)"
	@echo "  make migrate  - Run database migrations"
	@echo "  make dev      - Start database, run migrations, and run app"
	@echo "  make help     - Show this help"

db-start: ## Start database
	docker-compose up -d

db-stop: ## Stop database
	docker-compose down
run: ## Run the app with air (hot reload)
	air

migrate: ## Run database migrations
	cd internal/store/pgstore/migrations && \
	export $$(cat ../../../../.env | xargs) && \
	tern migrate --config tern.conf

dev: db-start migrate run ## Start database, run migrations, and run app

