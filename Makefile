# Определяем переменные (подгружаем из .env если нужно)
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Переменные для удобства
APP_NAME=warehouse-api
DOCKER_COMPOSE=docker-compose.yml

.PHONY: build run test migrate-up migrate-down docker-up docker-down docker-rebuild

# Локальная работа
build:
	go build -o bin/$(APP_NAME) ./cmd/api/main.go

run:
	go run ./cmd/api/main.go

test:
	go test -v ./internal/usecase/... ./internal/repository/...

# Работа с Docker
docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-rebuild:
	docker compose up --build

# Миграции (через CLI golang-migrate, если установлен)
# DB_URL берется из .env. Важно: для локального запуска адрес может быть localhost:5432
migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

# Создание новой миграции. Пример: make migrate-create name=add_users_table
migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)