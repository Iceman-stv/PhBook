package server

import (
	"PhBook/userCase"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// StartServer запускает локальный сервер
func StartServer(phoneBook *userCase.PhoneBook, port string) {
	// Обработчик для статических файлов
	fs := http.FileServer(http.Dir("PhBook/static"))
	http.Handle("/", fs)

	// Обработчик для получения контактов
	http.HandleFunc("/api/contacts", func(w http.ResponseWriter, r *http.Request) {
		// Получение userID из запроса
		userIDStr := r.URL.Query().Get("user_id")
		userID, err := strconv.Atoi(userIDStr)

		if err != nil {

			http.Error(w, "Неверный userID", http.StatusBadRequest)
			return
		}

		// Получение контактов для указанного userID
		contacts, err := phoneBook.GetContacts(userID)

		if err != nil {

			http.Error(w, "Ошибка при получении контактов", http.StatusInternalServerError)
			return
		}

		// Возвращение контактов в формате JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contacts)
	})

	fmt.Printf("Сервер запущен на http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
