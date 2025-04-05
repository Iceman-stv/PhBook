package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Logger определяет интерфейс для логгера
type Logger interface {
	LogInfo(message string, args ...interface{})
	LogError(message string, args ...interface{})
	LogFatal(message string, args ...interface{})
	LogWarn(message string, args ...interface{})
	Close() error
	Sync() error
}

// DefaultLogger реализует Logger с использованием стандартного логгера
type DefaultLogger struct {
	logFile *os.File
	done    chan struct{}
	mu      sync.Mutex
}

// InitLogger инициализирует и возвращает логгер
func InitLogger(logDir string) (Logger, error) {
	// Создание папки для логов, если она не существует
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {

		return nil, fmt.Errorf("ошибка при создании папки для логов: %v", err)
	}

	// Создание файла для логов
	logFile, err := os.OpenFile(filepath.Join(logDir, "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {

		return nil, fmt.Errorf("ошибка при открытии файла логов: %v", err)
	}

	logger := &DefaultLogger{
		logFile: logFile,
		done:    make(chan struct{}),
	}

	// Настройка глобального логгера
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)

	// Горутина для периодической очистки старых логов
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := cleanOldLogs(logDir, 7*24*time.Hour); err != nil {

					log.Printf("ошибка при очистке логов: %v", err)
				}
			case <-logger.done:
				return
			}
		}
	}()

	return logger, nil
}

func (l *DefaultLogger) LogInfo(message string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	log.Printf("[INFO] "+message, args...)
}

func (l *DefaultLogger) LogError(message string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	log.Printf("[ERROR] "+message, args...)
}

func (l *DefaultLogger) LogFatal(message string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	log.Fatalf("[FATAL] "+message, args...)
}

func (l *DefaultLogger) LogWarn(message string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	log.Printf("[WARNING] "+message, args...)
}

func (l *DefaultLogger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.done != nil {

		close(l.done)
		l.done = nil
	}

	if l.logFile != nil {

		err := l.logFile.Close()
		l.logFile = nil
		return err
	}
	return nil
}

func (l *DefaultLogger) Sync() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.logFile != nil {

		return l.logFile.Sync()
	}
	return nil
}

func cleanOldLogs(logDir string, maxAge time.Duration) error {
	files, err := ioutil.ReadDir(logDir)
	if err != nil {

		return fmt.Errorf("ошибка при чтении папки логов: %v", err)
	}

	for _, file := range files {
		if file.Name() == "app.log" {

			continue // Текущий лог-файл не удаляется
		}

		if time.Since(file.ModTime()) > maxAge {

			filePath := filepath.Join(logDir, file.Name())
			if err := os.Remove(filePath); err != nil {

				return fmt.Errorf("ошибка при удалении файла %s: %v", filePath, err)
			}
		}
	}
	return nil
}
