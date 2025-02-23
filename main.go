// PhBook
package main

import (
	"PhBook/database"
	"PhBook/interface/console"
	"PhBook/logger"
	"PhBook/userCase"
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

	//Старт консольного приложения
	app.Start()
}
