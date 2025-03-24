include .env

.PHONY: run build up reset db-status

run: build
	@./bin/${APP_NAME}

build:
	@go build -o ./bin/${APP_NAME} ./cmd/api/main.go

db-status:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} goose --dir=${MIGRATIONS_DIR} status

up:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} goose --dir=${MIGRATIONS_DIR} up

reset: 
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} goose --dir=${MIGRATIONS_DIR} reset

make-migration:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING} goose --dir=${MIGRATIONS_DIR} create ${name} sql
