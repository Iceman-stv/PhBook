package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"PhBook/domen"
	"PhBook/interface/netInterface/v1"
	"PhBook/logger"
	"PhBook/server/middleware"
	"PhBook/userCase"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server представляет HTTP-сервер приложения
type Server struct {
	router     *mux.Router         // Маршрутизатор запросов
	pb         *userCase.PhoneBook // Бизнес-логика
	logger     logger.Logger       // Логгер
	httpServer *http.Server        // HTTP-сервер для graceful shutdown
}

// NewServer создает новый экземпляр сервера
func NewServer(pb *userCase.PhoneBook, logger logger.Logger) *Server {
	return &Server{
		router: mux.NewRouter(),
		pb:     pb,
		logger: logger,
	}
}

// configureRouter настраивает все маршруты и middleware
func (s *Server) configureRouter() {
	// Middleware для логирования всех запросов
	s.router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s.logger.LogInfo("Incoming request: %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	// Инициализация обработчиков
	authHandlers := netInterface.NewAuthHandlers(s.pb)
	contactHandlers := netInterface.NewContactHandlers(s.pb)

	// Публичные маршруты (без аутентификации)
	s.router.HandleFunc("/register", authHandlers.HandleRegister).Methods("POST")
	s.router.HandleFunc("/auth", authHandlers.HandleAuth).Methods("POST")

	// Приватные маршруты (требуют аутентификации)
	api := s.router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware(s.logger)) // Добавляем middleware аутентификации

	api.HandleFunc("/contacts", contactHandlers.HandleAddContact).Methods("POST")
	api.HandleFunc("/contacts", contactHandlers.HandleGetContacts).Methods("GET")
	api.HandleFunc("/contacts/{name}", contactHandlers.HandleDeleteContact).Methods("DELETE")
	api.HandleFunc("/contacts/search", contactHandlers.HandleFindContact).Methods("GET")
}

// Start запускает сервер с поддержкой graceful shutdown
func (s *Server) Start() error {
	s.configureRouter()

	// Настройка CORS
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	// Создание HTTP-сервера
	s.httpServer = &http.Server{
		Addr:    ":8080",
		Handler: cors(s.router),
	}

	// Канал для ошибок сервера
	serverErr := make(chan error, 1)

	// Запуск сервера в отдельной горутине
	go func() {
		s.logger.LogInfo("Сервер запущен на :8080")
		fmt.Println("Сервер запущен на :8080")
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {

			serverErr <- err
		}
	}()

	// Ожидание сигналов завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Блокировка до получения сигнала или ошибки сервера
	select {
	case err := <-serverErr:
		return err
	case <-quit:
		s.logger.LogInfo("Остановка сервера...")

		// Настройка контекста с таймаутом для graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Попытка для graceful shutdown
		if err := s.httpServer.Shutdown(ctx); err != nil {

			s.logger.LogError("Ошибка остановки сервера: %v", err)
			return err
		}

		s.logger.LogInfo("Сервер успешно остановлен")
		return nil
	}
}

// handleError обрабатывает ошибки и отправляет соответствующие HTTP-ответы
func (s *Server) handleError(w http.ResponseWriter, err error) {
	switch err {
	case domen.ErrUserExists:
		http.Error(w, "User already exists", http.StatusConflict)
	case domen.ErrInvalidCredentials:
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	case domen.ErrContactNotFound:
		http.Error(w, "Contact not found", http.StatusNotFound)
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	s.logger.LogError("Request failed: %v", err)
}

func (s *Server) GetHTTPServer() *http.Server {
	return s.httpServer
}
