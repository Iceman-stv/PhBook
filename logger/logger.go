// logger.go
package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

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

	if err != nil {

		ErrorLogger.Printf("%s: %v\n", message, err)
	} else {
		ErrorLogger.Println(message)
	}
}
