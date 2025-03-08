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
	//Инициализация БД
	db, err := database.NewSQLiteDB()

	if err != nil {

		panic("Ошибка при инициализации БД" + err.Error())
		return
	}

	//Инициализация логгера
	logDir := "logs" //Папка для логгера

	if err := logger.InitLogger(logDir); err != nil {

		panic("Ошибка инициализации логгера " + err.Error())
	}

	//Создание PhoneBook
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
