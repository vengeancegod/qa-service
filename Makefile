DB_DSN = host=localhost user=root password=root dbname=db sslmode=disable
MIGRATIONS_DIR = migrations

migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" down

migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" status

db-connect:
	docker compose exec postgres psql -U root -d db

db-tables:
	docker compose exec postgres psql -U root -d db -c "\dt"

test:
	go test -v ./...
