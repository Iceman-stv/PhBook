# Имя исполняемого файла
BINARY_NAME=PhBook
# Команда Go
GO=go

# Пути для генерации кода
PROTO_PATH=proto
GEN_PATH=gen
PROTO_FILE=PhBook.proto

# Цель по умолчанию (выполняется при вызове `make`)
all: build

# Сборка проекта
build:
	$(GO) build -o $(BINARY_NAME) .

# Запуск проекта
run:
	$(GO) run .

# Очистка скомпилированных файлов
clean:
	$(GO) clean
	rm -f $(BINARY_NAME)
	rm -rf $(GEN_PATH)/*

# Установка зависимостей
deps:
	$(GO) mod download

# Форматирование кода
fmt:
	$(GO) fmt ./...

# Проверка стиля кода
vet:
	$(GO) vet ./...

# Линтинг кода
lint:
	golint ./...

# Запуск тестов
test: unit integration

unit:
	$(GO) test -v -short ./...

integration:
	$(GO) test -v -tags=integration ./...

bench:
	$(GO) test -bench=. ./...

# Статический анализ
staticcheck:
	staticcheck ./...

# Все проверки: тесты, форматирование, линтинг и статический анализ
	check: test fmt vet lint staticcheck

# Генерация кода из .proto файлов
generate:
	mkdir -p $(GEN_PATH)
	protoc --go_out=$(GEN_PATH) --go-grpc_out=$(GEN_PATH) --proto_path=$(PROTO_PATH) $(PROTO_PATH)/$(PROTO_FILE)

# Флаг .PHONY для указания целей, которые не являются файлами
.PHONY: all build run clean deps fmt vet lint test staticcheck check generate