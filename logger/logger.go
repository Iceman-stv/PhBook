package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Logger определяет интерфейс для логгера
type Logger interface {
	LogInfo(message string, args ...interface{})
	LogError(message string, args ...interface{})
	LogFatal(message string, args ...interface{})
	LogWarn(message string, args ...interface{})
	Close() error
}

// DefaultLogger реализует Logger с использованием стандартного логгера
type DefaultLogger struct {
	logFile *os.File
	done    chan struct{}
}

// LogInfo логирует информационные сообщения
func (l *DefaultLogger) LogInfo(message string, args ...interface{}) {
	log.Printf("[INFO] "+message, args...)
}

// LogError логирует сообщения об ошибках
func (l *DefaultLogger) LogError(message string, args ...interface{}) {
	log.Printf("[ERROR] "+message, args...)
}

// LogFatal логирует сообщения для критических ошибок
func (l *DefaultLogger) LogFatal(message string, args ...interface{}) {
	log.Fatalf("[FATAL] "+message, args...)
}

// LogFatal логирует сообщения для предупреждений
func (l *DefaultLogger) LogWarn(message string, args ...interface{}) {
	log.Printf("[WARNING] "+message, args...)
}

// Функция закрытия логгера
func (l *DefaultLogger) Close() error {
	if l.done != nil {

		close(l.done)
	}
	if l.logFile != nil {

		return l.logFile.Close()
	}
	return nil
}

// Удаление логов, которым более одной недели
func cleanOldLogs(logDir string, maxAge time.Duration) error {
	files, err := ioutil.ReadDir(logDir)
	if err != nil {

		return fmt.Errorf("Ошибка при чтении папки логов: %v", err)
	}

	for _, file := range files {
		if time.Since(file.ModTime()) > maxAge {

			filePath := filepath.Join(logDir, file.Name())
			if err := os.Remove(filePath); err != nil {

				return fmt.Errorf("Ошибка при удалении файла %s: %v", filePath, err)
			}
		}
	}

	return nil
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

	// Настройка логгера для записи в файл и определение формата
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetOutput(logFile)

	logger := &DefaultLogger{
		logFile: logFile,
		done:    make(chan struct{}),
	}

	// Горутина для переодической отчистки файла логов
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// Очистка логов старше 7ми дней
				if err := cleanOldLogs(logDir, 7*24*time.Hour); err != nil {
					log.Printf("Ошибка при очистке логов: %v", err)
				}
				// Завершение работы горутины
			case <-logger.done:
				return
			}
		}
	}()

	return logger, nil
}
