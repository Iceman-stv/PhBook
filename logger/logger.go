package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

<<<<<<< HEAD
var (
	ErrorLogger *log.Logger
)

// Инициализация логгера с определенной папкой для логов
func InitLogger(logDir string) error {
	//Создание папки для логов, если её нет

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {

		return fmt.Errorf("Ошибка при создании папки для логов %v", err)
	}

	//Файл для логов
	LogFile := filepath.Join(logDir, "errors.log")
	file, err := os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {

		log.Fatal("Ошибка при открытии файла логов", err)
	}
	//Инициализация логгера
	ErrorLogger = log.New(file, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

// функция для логгирования ошибок
func LogError(message string, err error) {
=======
// Logger определяет интерфейс для логгера
type Logger interface {
	LogInfo(message string, args ...interface{})
	LogError(message string, args ...interface{})
}

// DefaultLogger реализует Logger с использованием стандартного логгера
type DefaultLogger struct{}
>>>>>>> dop

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
	logDir = "logs" // Папка для логгера
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
