include .env

build:
	@go build -o ./.bin/tasks ./cmd/tasks

run:build
	@./.bin/tasks

up:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -dir ./sql/migrations postgres $(POSTGRES_DSN) up

down:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -dir ./sql/migrations postgres $(POSTGRES_DSN) down


reset:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -dir ./sql/migrations postgres $(POSTGRES_DSN) reset 

sqlc:
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@sqlc generate

tests:
	@POSTGRES_DSN=$(POSTGRES_DSN) go test -v ./...
