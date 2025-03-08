package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"PhBook/userCase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

// Server структура для хранения зависимостей сервера
type Server struct {
	router *mux.Router
	pb     *userCase.PhoneBook
}

// Конфигурация JWT
var jwtSecret = []byte("Qwert531") // Cекретный ключ

// Claims — структура для хранения данных в JWT
type Claims struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}

// NewServer создает новый экземпляр сервера
func NewServer(pb *userCase.PhoneBook) *Server {
	s := &Server{
		router: mux.NewRouter(),
		pb:     pb,
	}
	s.configureRouter()
	return s
}

// configureRouter настраивает маршруты сервера
func (s *Server) configureRouter() {
	// Открытые маршруты
	s.router.HandleFunc("/register", s.handleRegister()).Methods("POST")
	s.router.HandleFunc("/auth", s.handleAuth()).Methods("POST")

	// Защищенные маршруты
	api := s.router.PathPrefix("/api").Subrouter()
	api.Use(AuthMiddleware)
	api.HandleFunc("/contacts", s.handleAddContact()).Methods("POST")
	api.HandleFunc("/contacts", s.handleGetContacts()).Methods("GET")
	api.HandleFunc("/contacts/{name}", s.handleDeleteContact()).Methods("DELETE")
	api.HandleFunc("/contacts/search", s.handleFindContact()).Methods("GET")
}

// Запуск сервера
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// GenerateJWT создает новый JWT для пользователя
func GenerateJWT(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Токен действителен 24 часа

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT проверяет JWT и возвращает claims
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {

		return nil, err
	}

	if !token.Valid {

		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

// AuthMiddleware проверяет JWT и добавляет userID в контекст запроса
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {

			http.Error(w, "Ошибка в заголовке запроса", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {

			http.Error(w, "Неправильный формат токена", http.StatusUnauthorized)
			return
		}

		claims, err := ValidateJWT(tokenString)
		if err != nil {

			http.Error(w, "Неправильный токен", http.StatusUnauthorized)
			return
		}

		// Добавление userID в контекст запроса
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// Обработчики
func (s *Server) handleRegister() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

			http.Error(w, "Неправильный запрос (регистрация)", http.StatusBadRequest)
			return
		}

		if err := s.pb.RegisterUser(req.Username, req.Password); err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Пользователь зарегистрирован"))
	}
}

func (s *Server) handleAuth() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

			http.Error(w, "Неправильный запрос (Аутентификация)", http.StatusBadRequest)
			return
		}

		userID, err := s.pb.AuthUser(req.Username, req.Password)
		if err != nil {

			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := GenerateJWT(userID)
		if err != nil {

			http.Error(w, "Ошибка при генерации токена", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}

func (s *Server) handleAddContact() http.HandlerFunc {
	type request struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

			http.Error(w, "Неправильный запрос (Добавление контакта)", http.StatusBadRequest)
			return
		}

		userID := r.Context().Value("userID").(int)
		if err := s.pb.AddContact(userID, req.Name, req.Phone); err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Контакт добавлен"))
	}
}

func (s *Server) handleGetContacts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID").(int)
		contacts, err := s.pb.GetContacts(userID)
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(contacts)
	}
}

func (s *Server) handleDeleteContact() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID").(int)
		name := mux.Vars(r)["name"]
		if err := s.pb.DelContact(userID, name); err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Контакт удалён"))
	}
}

func (s *Server) handleFindContact() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID").(int)
		name := r.URL.Query().Get("name")

		contacts, err := s.pb.FindContact(userID, name)
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(contacts)
	}
}
