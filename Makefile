# PlanPal backend – common commands
# Install migrate: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: run migrate-up migrate-down migrate-create tidy

run:
	go run ./cmd/api

tidy:
	go mod tidy

# Migrations (requires DATABASE_URL or default from .env)
migrate-up:
	migrate -path db/migrations -database "$${DATABASE_URL:-postgres://planpal:planpal@localhost:5432/planpal?sslmode=disable}" up

migrate-down:
	migrate -path db/migrations -database "$${DATABASE_URL:-postgres://planpal:planpal@localhost:5432/planpal?sslmode=disable}" down 1

# Create a new migration: make migrate-create NAME=add_avatar
migrate-create:
	@name=$${NAME:-new_migration}; migrate create -ext sql -dir db/migrations -seq $$name
