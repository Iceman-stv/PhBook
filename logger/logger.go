package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Logger определяет интерфейс для логгера
type Logger interface {
	LogInfo(message string, args ...interface{})
	LogError(message string, args ...interface{})
}

// DefaultLogger реализует Logger с использованием стандартного логгера
type DefaultLogger struct{}

// LogInfo логирует информационные сообщения
func (l *DefaultLogger) LogInfo(message string, args ...interface{}) {
	log.Printf("[INFO] "+message, args...)
}

// LogError логирует сообщения об ошибках
func (l *DefaultLogger) LogError(message string, args ...interface{}) {
	log.Printf("[ERROR] "+message, args...)
}

// InitLogger инициализирует и возвращает логгер
func InitLogger(logDir string) (Logger, error) {
	// Создание папки для логов, если она не существует
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("Ошибка при создании папки для логов: %v", err)
	}

	// Создание файла для логов
	logFile, err := os.OpenFile(filepath.Join(logDir, "app.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при открытии файла логов: %v", err)
	}

	// Настройка логгера для записи в файл
	log.SetOutput(logFile)

	return &DefaultLogger{}, nil
}
