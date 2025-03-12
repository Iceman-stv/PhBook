// PhBook
package main

import (
	"PhBook/database"
	"PhBook/interface/console"
	"PhBook/logger"
	"PhBook/server"
	"PhBook/userCase"
	"time"
)

func main() {
	// Инициализация логгера
	l, err := logger.InitLogger("Logs")
	if err != nil {

		panic("Ошибка инициализации логгера: " + err.Error())
	}

	// Инициализация БД
	db, err := database.NewSQLiteDB(l)
	if err != nil {

		l.LogError("Ошибка при инициализации БД: %v", err)
		return
	}

	// Создание PhoneBook
	pb := userCase.NewPhoneBook(db)

	//Создание консольного приложения
	app := console.NewConsole(pb)

	// Запуск локального сервера
	go func() {
		serv := server.NewServer(pb, l)
		serv.Start()
	}()

	//Старт консольного приложения
	app.Start()

	time.Sleep(5 * time.Second)
}
