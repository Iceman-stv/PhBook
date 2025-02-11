// PhBook
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	
)

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

func main() {
	phoneBook := PhoneBook{}
	filename := "contacts.csv"
	err := phoneBook.LoadContacts(filename)
	if err != nil {
		fmt.Println("Ошибка при загрузке контактов:", err)
	}

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
			err := phoneBook.SaveContacts(filename)
			if err != nil {
				fmt.Println("Ошибка при сохранении контактов:", err)
			}
		case 6:
			err := phoneBook.LoadContacts(filename)
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
}
