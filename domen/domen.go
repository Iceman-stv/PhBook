package domen

// User представляет структуру пользователя
type User struct {
	ID       int    `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Contact представляет структуру контакта
type Contact struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID int    `json:"userID"`
}
