# ==================== Этап сборки ====================
FROM golang:1.22-alpine AS builder

# Установка репозиториев Alpine
RUN echo "https://dl-cdn.alpinelinux.org/alpine/latest-stable/main" > /etc/apk/repositories && \
    echo "https://dl-cdn.alpinelinux.org/alpine/latest-stable/community" >> /etc/apk/repositories && \
    apk update

WORKDIR /app

# Установка зависимостей для сборки
RUN apk add --no-cache --virtual .build-deps \
    git \
    make \
    gcc \
    musl-dev \
    sqlite-dev \
    protobuf-dev

# Установка инструментов
RUN go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Копирование зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копирование исходного кода
COPY . .

# Генерация gRPC кода (если proto-файл существует)
RUN if [ -f "proto/PhBook.proto" ]; then \
      mkdir -p gen && \
      protoc --go_out=gen --go-grpc_out=gen \
        --proto_path=proto proto/PhBook.proto; \
    fi

# Бинарник
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o /phbook .

# ==================== Финальный образ ====================
FROM alpine:latest

# Обновление репозиториев и установка runtime-зависимостей
RUN echo "https://dl-cdn.alpinelinux.org/alpine/latest-stable/main" > /etc/apk/repositories && \
    echo "https://dl-cdn.alpinelinux.org/alpine/latest-stable/community" >> /etc/apk/repositories && \
    apk update && \
    apk add --no-cache \
    ca-certificates \
    sqlite \
    libc6-compat && \
    adduser -D -u 1000 appuser && \
    rm -rf /var/cache/apk/*

# Копирование
COPY --from=builder /phbook /app/phbook
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Настройка прав
RUN chown -R appuser:appuser /app && \
    chmod 550 /app/phbook && \
    chmod 750 /app/migrations && \
    mkdir -p /app/logs && \
    chown appuser:appuser /app/logs

# Рабочая директория и пользователь
USER appuser
WORKDIR /app

# Точка входа
CMD ["sh", "-c", "migrate -path=/app/migrations -database='sqlite3:///app/phonebook.db?_busy_timeout=5000' up && /app/phbook --server"]