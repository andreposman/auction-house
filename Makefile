# Auction House API - Simple Makefile

.PHONY: start stop run dev migrate

start: ## Start database
	docker-compose up -d

stop: ## Stop database
	docker-compose down

run: ## Run the app
	go run cmd/api/main.go

migrate: ## Run database migrations
	cd internal/store/pgstore/migrations && tern migrate --config tern.conf

dev: start migrate run ## Start database, run migrations, and run app
