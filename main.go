// PhBook
package main

import (
	"PhBook/database"
	"PhBook/interface/console"
	"PhBook/logger"
	"PhBook/server"
	"PhBook/userCase"
)

func main() {
	// Инициализация логгера
	logger, err := logger.InitLogger("Logs")
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
		server := server.NewServer(pb, logger)
		server.Start()
	}()

	//Старт консольного приложения
	app.Start()
}
