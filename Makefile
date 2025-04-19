# --- Настройка shell для Windows/Unix ---
ifeq ($(OS),Windows_NT)
# Для Windows
	SHELL = cmd.exe
	.SHELLFLAGS = /c
	NULLDEV = nul
	TRUE = (exit 0)
else
# Для Linux/Mac
	SHELL = /bin/sh
	NULLDEV = /dev/null
	TRUE = true
endif

# --- Конфигурация ---
BINARY_NAME = PhBook
GO = go
PORT = 8080
DB_PATH = /app/phonebook.db
DOCKER_IMAGE = phbook
DOCKER_TAG = latest

# --- Автодетект ОС и настройка путей ---
ifeq ($(OS),Windows_NT)
# Windows
	CURDIR := $(subst \,/,${CURDIR})
	DOCKER_VOL :=
	DOCKER_PATHS = \
		-v "$(CURDIR)/migrations:/app/migrations:ro" \
		-v "$(CURDIR)/phonebook.db:/app/phonebook.db"
else
# Linux/Mac
	CURDIR := $(shell pwd)
	DOCKER_VOL :=
	DOCKER_PATHS = \
		-v "$(CURDIR)/migrations:/app/migrations:ro" \
		-v "$(CURDIR)/phonebook.db:/app/phonebook.db"
endif

# --- Основные цели ---
.PHONY: all build run clean test generate check-deps

all: build

build:
	$(GO) build -o $(BINARY_NAME) .

run:
	$(GO) run .

clean:
	@echo Cleaning in progress...
	$(GO) clean
	@if exist $(BINARY_NAME) del /Q $(BINARY_NAME)
	@docker stop $(DOCKER_IMAGE) 2>nul || echo Container is stopped
	@docker rm $(DOCKER_IMAGE) 2>nul || echo Container is delete
	@echo Cleaning success

test:
	$(GO) test -v ./...

generate:
	mkdir -p gen
	protoc --go_out=gen --go-grpc_out=gen --proto_path=proto proto/PhBook.proto

check-deps:  # Проверка зависимостей
	@protoc --version || (echo "err: protoc not installed"; exit 1)

# --- Docker-команды ---
.PHONY: docker-build docker-run docker-stop docker-logs docker-migrate docker-full

docker-build:  # Сборка Docker-образа
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-run:  # Запуск контейнера
	docker run -d \
	-p $(PORT):$(PORT) \
	-v "$(CURDIR)/logs:/app/logs" \
	-v "$(CURDIR)/phonebook.db:$(DB_PATH)" \
	-v "$(CURDIR)/migrations:/app/migrations" \
	--name $(DOCKER_IMAGE) \
	$(DOCKER_IMAGE):$(DOCKER_TAG)

docker-errors:  # Просмотр ошибок
	docker logs $(DOCKER_IMAGE) | grep -i "error\|warning"

docker-stop:  # Остановка контейнера
	@docker stop $(DOCKER_IMAGE) 2>$(NULLDEV) || $(TRUE)
	@docker rm $(DOCKER_IMAGE) 2>$(NULLDEV) || $(TRUE)

docker-logs:  # Просмотр логов
	docker logs -f $(DOCKER_IMAGE)

docker-migrate:  # Запуск миграций
	docker run --rm \
	$(DOCKER_PATHS) \
	$(DOCKER_IMAGE):$(DOCKER_TAG) \
	migrate -path=/app/migrations -database="sqlite3:///app/phonebook.db?_busy_timeout=5000" up

docker-full: check-deps docker-build docker-migrate docker-run  # Полный цикл

# --- Утилиты ---
.PHONY: db-backup

db-backup:
	cp phonebook.db phonebook_$(shell date +%Y%m%d).db

status:
	@echo "=============================="
	@echo "= Project: PhBook"
	@echo "= Go ver.: $(shell go version)"
	@echo "= Docker: $(shell docker --version)"
	@echo "=============================="
