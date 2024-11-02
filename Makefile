build:
	@go build -o ./.bin/tasks ./cmd/tasks

run:build
	@./.bin/tasks

up:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -d sql/migrations postgres $(POSTGRES_DSN) up

down:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -d sql/migrations postgres $(POSTGRES_DSN) down
