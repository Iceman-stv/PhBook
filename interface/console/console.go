// console
package console

import (
	"PhBook/userCase"
	"fmt"
)

// Интерфейс для взаимодействия с пользователем
type Console struct {
	phoneBook *userCase.PhoneBook
}

// Создание нового экземпляра консоли
func NewConsole(phoneBook *userCase.PhoneBook) *Console {
	return &Console{phoneBook: phoneBook}
}

// Запуск консоли
func (c *Console) Start() {
	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Регистрация")
		fmt.Println("2. Войти")
		fmt.Println("3. Выйти")

		var choice int
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			c.registerUser()

		case 2:
			c.authUser()

		case 3:

			fmt.Println("Выход из программы.")
			return

		default:

			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func (c *Console) registerUser() {
	var username, password string

	fmt.Print("Введите имя: ")
	fmt.Scanln(&username)
	fmt.Print("Введите пароль: ")
	fmt.Scanln(&password)

	err := c.phoneBook.RegisterUser(username, password)

	if err != nil {

		fmt.Println("Ошибка при регистрации пользователя")
	} else {
		fmt.Println("Пользователь зарегистрирован")
	}
}

func (c *Console) authUser() {
	var username, password string

	fmt.Print("Введите имя: ")
	fmt.Scanln(&username)
	fmt.Print("Введите пароль: ")
	fmt.Scanln(&password)

	userID, err := c.phoneBook.AuthUser(username, password)

	if err != nil {

		fmt.Println("Ошибка аутентификации")
	} else {
		fmt.Println("Добро пожаловать", username)
		c.phoneBookMenu(userID)
	}
}

// Меню телефонной книги
func (c *Console) phoneBookMenu(userID int) {
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

			c.addContact(userID)

		case 2:

			c.delContact(userID)

		case 3:

			c.findContact(userID)

		case 4:

			c.getContacts(userID)

		case 5:

			fmt.Println("Выход из программы.")
			return

		default:

			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func (c *Console) addContact(userID int) {
	var name, phone string

	fmt.Print("Введите имя: ")
	fmt.Scanln(&name)
	fmt.Print("Введите телефон: ")
	fmt.Scanln(&phone)

	err := c.phoneBook.AddContact(userID, name, phone)

	if err != nil {

		fmt.Println("Ошибка, контакт не добавлен")
	} else {
		fmt.Println("Контакт добавлен")
	}
}

func (c *Console) delContact(userID int) {
	var name string

	fmt.Print("Введите имя для удаления: ")
	fmt.Scanln(&name)

	err := c.phoneBook.DelContact(userID, name)

	if err != nil {

		fmt.Println("Ошибка, контакт не удален")
	} else {
		fmt.Println("Контакт удален")
	}
}

func (c *Console) findContact(userID int) {
	var name string

	fmt.Print("Введите имя для поиска: ")
	fmt.Scanln(&name)

	contacts, err := c.phoneBook.FindContact(userID, name)

	if err != nil || contacts == nil {

		fmt.Println("Ошибка, контакты не найдены")
	} else {
		fmt.Println("Найденные контакты:")

		for _, contact := range contacts {
			fmt.Printf("%s - %s\n", contact.Name, contact.Phone)
		}
	}
}

func (c *Console) getContacts(userID int) {

	contacts, err := c.phoneBook.GetContacts(userID)

	if err != nil {

		fmt.Println("Ошибка, невозможно вывести контакты")
	} else {
		fmt.Println("Kонтакты:")

		for _, contact := range contacts {
			fmt.Printf("%s - %s\n", contact.Name, contact.Phone)
		}
	}
}
