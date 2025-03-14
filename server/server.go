package server

import (
	"fmt"
	"log"
	"net/http"

	"PhBook/interface/user_http/v1"
	"PhBook/logger"
	"PhBook/server/middleware"
	"PhBook/userCase"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	pb     *userCase.PhoneBook
	logger logger.Logger
}

func NewServer(pb *userCase.PhoneBook, logger logger.Logger) *Server {
	s := &Server{
		router: mux.NewRouter(),
		pb:     pb,
		logger: logger,
	}
	s.configureRouter()
	return s
}

func (s *Server) configureRouter() {
	// Регистрация обработчиков
	authHandlers := user_http.NewAuthHandlers(s.pb)
	contactHandlers := user_http.NewContactHandlers(s.pb)

	s.router.HandleFunc("/register", authHandlers.HandleRegister).Methods("POST")
	s.router.HandleFunc("/auth", authHandlers.HandleAuth).Methods("POST")

	api := s.router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware(s.logger))
	api.HandleFunc("/contacts", contactHandlers.HandleAddContact).Methods("POST")
	api.HandleFunc("/contacts", contactHandlers.HandleGetContacts).Methods("GET")
	api.HandleFunc("/contacts/{name}", contactHandlers.HandleDeleteContact).Methods("DELETE")
	api.HandleFunc("/contacts/search", contactHandlers.HandleFindContact).Methods("GET")
}

func (s *Server) Start() {
	s.logger.LogInfo("Сервер заущен на :8080")
	fmt.Println("Сервер заущен на :8080")
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(s.router)))
}
