// PhBook
package main

import (
	"PhBook/database"
	"PhBook/interface/console"
	"PhBook/userCase"
	"fmt"
)

func main() {
	//Инициализация БД
	db, err := database.NewSQLiteDB()

	if err != nil {

		fmt.Printf("Ошибка при инициализации БД %v", err)
		return
	}

	//Создание PhoneBook
	pb := userCase.NewPhoneBook(db)

	//Создание консольного приложения
	app := console.NewConsole(pb)

	//Старт консольного приложения
	app.Start()
}
