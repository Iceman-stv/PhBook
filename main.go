package main

import (
	"PhBook/database"
	"PhBook/interface/console"
	"PhBook/logger"
	"PhBook/server"
	"PhBook/server/gRPC"
	"PhBook/userCase"
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func runServerMode(l logger.Logger, pb *userCase.PhoneBook) {
	errChan := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	// HTTP сервер
	httpServer := server.NewServer(pb, l)
	go func() {
		defer wg.Done()
		l.LogInfo("HTTP сервер запущен на :8080")
		if err := httpServer.Start(); err != nil && err != http.ErrServerClosed {

			errChan <- err
		}
	}()

	// gRPC сервер
	grpcServer := gRPC.New(pb, l, gRPC.DefaultConfig(l))
	go func() {
		defer wg.Done()
		l.LogInfo("gRPC сервер запущен на :50051")
		if err := grpcServer.Start(); err != nil {

			errChan <- err
		}
	}()

	// Ожидание завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		l.LogFatal("Ошибка сервера: %v", err)
	case <-quit:
		l.LogInfo("Получен сигнал завершения...")
	}

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	grpcServer.Stop()
	if srv := httpServer.GetHTTPServer(); srv != nil {

		srv.Shutdown(ctx)
	}
	wg.Wait()
	l.LogInfo("Все серверы остановлены. Приложение завершено.")
}

func main() {
	// Инициализация
	l, err := logger.InitLogger("Logs")
	if err != nil {

		panic("Ошибка логгера: " + err.Error())
	}
	defer l.Close()

	db, err := database.NewSQLiteDB(l)
	if err != nil {

		l.LogFatal("Ошибка БД: %v", err)
	}
	pb := userCase.NewPhoneBook(db)

	// Выбор режима
	if len(os.Args) > 1 && os.Args[1] == "--server" {

		l.LogInfo("=== РЕЖИМ СЕРВЕРА ===")
		runServerMode(l, pb)
	} else {

		l.LogInfo("=== КОНСОЛЬНЫЙ РЕЖИМ ===")
		console.NewConsole(pb).Start()
	}
}
