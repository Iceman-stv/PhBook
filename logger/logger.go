// logger.go
package logger

import (
	"log"
	"os"
)

var (
	ErrorLogger *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {

		log.Fatal("Ошибка при открытии файла логов", err)
	}

	ErrorLogger = log.New(file, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogError(message string, err error) {

	if err != nil {

		ErrorLogger.Printf("%s: %v\n", message, err)
	} else {
		ErrorLogger.Println(message)
	}
}
