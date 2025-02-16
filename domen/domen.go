package domen

// User представляет структуру пользователя
type User struct {
	ID       int
	Username string
	Password string
}

// Contact представляет структуру контакта
type Contact struct {
	ID     int
	Name   string
	Phone  string
	UserID int
}