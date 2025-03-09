// PhBook
package main

import (
	"PhBook/database"
	"PhBook/interface/console"
	"PhBook/logger"
	"PhBook/server"
	"PhBook/userCase"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	// Инициализация логгера
	logDir := "logs" // Папка для логгера
	logger, err := logger.InitLogger(logDir)
	if err != nil {
		
		panic("Ошибка инициализации логгера: " + err.Error())
	}

	// Инициализация БД
	db, err := database.NewSQLiteDB(logger)
	if err != nil {

		logger.LogError("Ошибка при инициализации БД: %v", err)
		return
	}

	// Создание PhoneBook
	pb := userCase.NewPhoneBook(db)

	//Создание консольного приложения
	app := console.NewConsole(pb)

	// Запуск локального сервера
	go func() {
		server := server.NewServer(pb)

		// Найстройка Cross-origin-resource-sharing (CORS)
		headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
		methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})
		origins := handlers.AllowedOrigins([]string{"*"})
		// Запуск сервера
		log.Println("Сервер запущен на адресе :8080")
		log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(server)))
	}()

	//Старт консольного приложения
	app.Start()
}
