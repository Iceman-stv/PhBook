// PhBook
package main

import (
	"PhBook/database"
	"PhBook/interface/console"
	"PhBook/logger"
<<<<<<< HEAD
=======
	"PhBook/server"
>>>>>>> dop
	"PhBook/userCase"
	"time"
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

<<<<<<< HEAD
	//Инициализация логгера
	logDir := "logs" //Папка для логгера

	if err := logger.InitLogger(logDir); err != nil {

		panic("Ошибка инициализации логгера " + err.Error())
	}

	//Создание PhoneBook
=======
	// Создание PhoneBook
>>>>>>> dop
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

	time.Sleep(5 * time.Second)
}
