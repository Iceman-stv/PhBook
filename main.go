// PhBook
package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// User представляет структуру пользователя

type User struct {
	Username string
	Password string
}

// Contact представляет структуру контакта
type Contact struct {
	Name  string
	Phone string
}

// PhoneBook представляет телефонную книгу
type PhoneBook struct {
	contacts []Contact
}

// AddContact добавляет новый контакт
func (pb *PhoneBook) AddContact(name, phone string) {
	pb.contacts = append(pb.contacts, Contact{Name: name, Phone: phone})
	fmt.Println("Контакт добавлен.")
}

// DeleteContact удаляет контакт по имени
func (pb *PhoneBook) DeleteContact(name string) {
	for i, contact := range pb.contacts {
		
		if contact.Name == name {

			pb.contacts = append(pb.contacts[:i], pb.contacts[i+1:]...)
			fmt.Println("Контакт удален.")
			return
		}
	}
	fmt.Println("Контакт не найден.")
}

// FindContact ищет контакт по имени
func (pb *PhoneBook) FindContact(name string) {
	for _, contact := range pb.contacts {
		
		if contact.Name == name {

			fmt.Printf("Найден контакт: %s - %s\n", contact.Name, contact.Phone)
			return
		}
	}
	fmt.Println("Контакт не найден.")
}

// ListContacts выводит все контакты
func (pb *PhoneBook) ListContacts() {
	if len(pb.contacts) == 0 {

		fmt.Println("Телефонная книга пуста.")
		return
	}
	fmt.Println("Телефонная книга:")

	for _, contact := range pb.contacts {
		fmt.Printf("%s - %s\n", contact.Name, contact.Phone)
	}
}

// SaveContacts сохраняет контакты в файл
func (pb *PhoneBook) SaveContacts(filename string) error {
	file, err := os.Create(filename)

	if err != nil {

		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, contact := range pb.contacts {
		err := writer.Write([]string{contact.Name, contact.Phone})
		if err != nil {

			return err
		}
	}
	fmt.Println("Контакты успешно сохранены в файл")
	return nil
}

// LoadContacts - загружает данные из файла
func (pb *PhoneBook) LoadContacts(filename string) error {
	file, err := os.Open(filename)
	if err != nil {

		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {

		return err
	}

	for _, record := range records {
		pb.contacts = append(pb.contacts, Contact{Name: record[0], Phone: record[1]})
	}
	fmt.Println("Контакты успешно загружены")
	return nil
}

// Хэширование пароля
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {

		return "", err
	}

	return string(hashedPassword), nil
}

// Проверка пароля
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

// Сохранение пользователя
func SaveUsers(filename string, users []User) error {
	file, err := os.Create(filename)

	if err != nil {

		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, user := range users {
		err := writer.Write([]string{user.Username, user.Password})

		if err != nil {

			return err
		}
	}

	fmt.Println("Пользователи сохранены")
	return nil
}

// Загрузка пользователя
func LoadUsers(filename string) ([]User, error) {
	file, err := os.Open(filename)
	if err != nil {

		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {

		return nil, err
	}

	var users []User

	for _, record := range records {
		users = append(users, User{Username: record[0], Password: record[1]})
	}
	fmt.Println("Загрузка пользователей прошла успешно")
	return users, nil
}

// Регистрация нового пользователя

func RegisterUser(users *[]User, username, password string) error {
	// Проверка существует ли пользователь

	for _, user := range *users {

		if user.Username == username {
			return fmt.Errorf("Пользователь уже зарегистрирован")
		}
	}

	//Хэш пароля
	hashedPassword, err := HashPassword(password)

	if err != nil {

		return err
	}

	//Добавление пользователя
	*users = append(*users, User{Username: username, Password: hashedPassword})
	fmt.Println("Пользователь зарегистрирован")
	return nil
}

//Аутентификация пользователя
func AuthUser(users []User, username, password string) bool {
	for _, user := range users {

		if user.Username == username && CheckPassword(user.Password, password) {

			return true
		}
	}
	return false
}

func main() {
	usersFilename := "users.csv"
	contactsFilename := "contacts.csv"

	// Загрузка пользователей
	users, err := LoadUsers(usersFilename)

	if err != nil && !os.IsNotExist(err) {

		fmt.Println("Ошибка при загрузке пользователя:", err)
		return
	}

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

			err := RegisterUser(&users, username, password)

			if err != nil {

				fmt.Println("Ошибка:", err)
			} else {
				err := SaveUsers(usersFilename, users)

				if err != nil {

					fmt.Println("Ошибка при сохранении:", err)
				}
			}

		case 2:

			var username, password string

			fmt.Print("Введите имя: ")
			fmt.Scanln(&username)
			fmt.Print("Введите пароль: ")
			fmt.Scanln(&password)

			if AuthUser(users, username, password) {

				fmt.Println("Добро пожаловать!", username)
				phoneBook := PhoneBook{}
				err := phoneBook.LoadContacts(contactsFilename)

				if err != nil {

					fmt.Println("Ошибка при загрузке контактов:", err)
				}

				//Меню телефонной книги
				for {
					fmt.Println("\nВыберите действие:")
					fmt.Println("1. Добавить контакт")
					fmt.Println("2. Удалить контакт")
					fmt.Println("3. Найти контакт")
					fmt.Println("4. Вывести все контакты")
					fmt.Println("5. Сохранить все контакты")
					fmt.Println("6. Загрузка контактов")
					fmt.Println("7. Выйти")

					var choice int
					fmt.Scanln(&choice)

					switch choice {

					case 1:

						var name, phone string

						fmt.Print("Введите имя: ")
						fmt.Scanln(&name)
						fmt.Print("Введите телефон: ")
						fmt.Scanln(&phone)
						phoneBook.AddContact(name, phone)

					case 2:

						var name string

						fmt.Print("Введите имя для удаления: ")
						fmt.Scanln(&name)
						phoneBook.DeleteContact(name)

					case 3:

						var name string

						fmt.Print("Введите имя для поиска: ")
						fmt.Scanln(&name)
						phoneBook.FindContact(name)

					case 4:

						phoneBook.ListContacts()

					case 5:

						err := phoneBook.SaveContacts(contactsFilename)

						if err != nil {

							fmt.Println("Ошибка при сохранении контактов:", err)
						}

					case 6:

						err := phoneBook.LoadContacts(contactsFilename)

						if err != nil {

							fmt.Println("Ошибка при загрузке контактов:", err)
						}

					case 7:

						fmt.Println("Выход из программы.")
						return

					default:

						fmt.Println("Неверный выбор. Попробуйте снова.")
					}
				}
			} else {
				fmt.Println("Ошибка аутентификации. Неверное имя пользователя или пароль")
			}

		case 3:
			fmt.Println("Выход из программы.")
			return

		default:

			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}
