# Подгружаем переменные из .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: help up down rebuild test logs

# Команда по умолчанию
help:
	@echo "Доступные команды:"
	@echo "  make up      - Запустить проект в Docker (в фоне)"
	@echo "  make down    - Остановить и удалить контейнеры"
	@echo "  make rebuild - Пересобрать образы и запустить проект"
	@echo "  make logs    - Посмотреть логи приложения"
	@echo "  make test    - Запустить unit-тесты (локально)"

# Запуск проекта
up:
	docker compose up -d

# Остановка проекта
down:
	docker compose down

# Полная пересборка и запуск (рекомендуется при изменении кода)
rebuild:
	docker compose up --build

# Просмотр логов приложения
logs:
	docker logs -f warehouse_app

# Запуск тестов
test:
	go test -v ./internal/usecase/...