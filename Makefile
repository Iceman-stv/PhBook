# Имя исполняемого файла
BINARY_NAME=PhBook
# Команда Go
GO=go

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
test:
	$(GO) test ./...

# Статический анализ
staticcheck:
	staticcheck ./...

# Все проверки: тесты, форматирование, линтинг и статический анализ
check: test fmt vet lint staticcheck

# Флаг .PHONY для указания целей, которые не являются файлами
.PHONY: all build run clean deps fmt vet lint test staticcheck check