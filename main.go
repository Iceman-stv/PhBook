// PhBook
package main

import (
	"PhBook/database"
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

	//Меню регистрации, входа
	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Регистрация")
		fmt.Println("2. Войти")
		fmt.Println("3. Выйти")

		var choice int
		fmt.Scanln(&choice)

		switch choice {

		case 1:

			var username, password string

			fmt.Print("Введите имя: ")
			fmt.Scanln(&username)
			fmt.Print("Введите пароль: ")
			fmt.Scanln(&password)

			err := pb.RegisterUser(username, password)

			if err != nil {

				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Пользователь зарегистрирован")
			}

		case 2:

			var username, password string

			fmt.Print("Введите имя: ")
			fmt.Scanln(&username)
			fmt.Print("Введите пароль: ")
			fmt.Scanln(&password)

			userID, err := pb.AuthUser(username, password)
			if err != nil {
				fmt.Println("Ошибка аутентификации:", err)
			} else {
				fmt.Println("Добро пожаловать", username)
				phoneBookMenu(pb, userID)
			}

		case 3:

			fmt.Println("Выход из программы.")
			return

		default:

			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

// Меню телефонной книги
func phoneBookMenu(pb *userCase.PhoneBook, userID int) {
	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Добавить контакт")
		fmt.Println("2. Удалить контакт")
		fmt.Println("3. Найти контакт")
		fmt.Println("4. Вывести все контакты")
		fmt.Println("5. Выйти")

		var choice int
		fmt.Scanln(&choice)

		switch choice {

		case 1:

			var name, phone string

			fmt.Print("Введите имя: ")
			fmt.Scanln(&name)
			fmt.Print("Введите телефон: ")
			fmt.Scanln(&phone)

			err := pb.AddContact(userID, name, phone)

			if err != nil {

				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Контакт добавлен")
			}

		case 2:

			var name string

			fmt.Print("Введите имя для удаления: ")
			fmt.Scanln(&name)

			err := pb.DelContact(userID, name)

			if err != nil {

				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Контакт удален")
			}

		case 3:

			var name string

			fmt.Print("Введите имя для поиска: ")
			fmt.Scanln(&name)

			contacts, err := pb.FindContact(userID, name)

			if err != nil {

				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Найденные контакты:")

				for _, contact := range contacts {
					fmt.Printf("%d: %s - %s\n", contact.ID, contact.Name, contact.Phone)
				}
			}

		case 4:

			contacts, err := pb.GetContacts(userID)

			if err != nil {

				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Kонтакты:")

				for _, contact := range contacts {
					fmt.Printf("%d: %s - %s\n", contact.ID, contact.Name, contact.Phone)
				}
			}

		case 5:

			fmt.Println("Выход из программы.")
			return

		default:

			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}
