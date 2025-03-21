// PhBook
package main

import (
	"PhBook/database"
	"PhBook/interface/console"
	"PhBook/logger"
	"PhBook/server"
	"PhBook/userCase"
	"PhBook/server/gRPC"
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
	
	// Запуск gRPC - сервера 
		go func() {
		servGRPC := gRPC.NewGRPCServer(pb, l)
		if err := servGRPC.Start(); err != nil {
			
			l.LogError("Ошибка при запуске gRPC-сервера %v", err)
		}
	}()


	//Старт консольного приложения
	app.Start()

	time.Sleep(5 * time.Second)
}
