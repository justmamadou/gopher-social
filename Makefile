include .envrc
MIGRATIONS_PATH=./cmd/migrate/migrations

.PHONY: migrate-create
migrate:
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DSN)" up

.PHONY: migrate-down
migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DSN)" down ${filter-out $@,$(MAKECMDGOALS)}