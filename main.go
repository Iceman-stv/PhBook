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

func main() {
	// Инициализация логгера
	l, err := logger.InitLogger("Logs")
	if err != nil {

		panic("Ошибка инициализации логгера: " + err.Error())
	}
	defer l.Close()

	// Инициализация БД
	db, err := database.NewSQLiteDB(l)
	if err != nil {

		l.LogFatal("Ошибка при инициализации БД: %v", err)
		return
	}

	// Создание PhoneBook
	pb := userCase.NewPhoneBook(db)

	// Канал для ошибок серверов
	errChan := make(chan error, 2)

	// Создание WaitGroup для ожидания завершения серверов
	var wg sync.WaitGroup
	wg.Add(2) // Для HTTP и gRPC серверов

	// Запуск HTTP-сервера
	httpServer := server.NewServer(pb, l)
	go func() {
		defer wg.Done()
		l.LogInfo("Запуск HTTP-сервера на :8080")
		if err := httpServer.Start(); err != nil && err != http.ErrServerClosed {

			errChan <- err
		}
	}()

	// Запуск gRPC-сервера
	grpcServer := gRPC.New(pb, l, gRPC.DefaultConfig())
	go func() {
		defer wg.Done()
		l.LogInfo("Запуск gRPC-сервера на :50051")
		if err := grpcServer.Start(); err != nil {

			errChan <- err
		}
	}()

	// Создание и запуск консольного приложения
	app := console.NewConsole(pb)
	go func() {
		app.Start()
	}()

	// Ожидание сигналов завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Ожидание либо сигнала завершения, либо ошибки от серверов
	select {
	case err := <-errChan:
		l.LogFatal("Ошибка сервера: %v", err)
	case <-quit:
		l.LogInfo("Получен сигнал завершения...")
	}

	// Graceful shutdown
	l.LogInfo("Завершение работы серверов...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Остановка gRPC сервера
	go func() {
		grpcServer.Stop()
	}()

	// Остановка HTTP сервера через его внутренний http.Server
	go func() {
		if srv := httpServer.GetHTTPServer(); srv != nil {

			if err := srv.Shutdown(ctx); err != nil {

				l.LogError("Ошибка при остановке HTTP-сервера: %v", err)
			}
		}
	}()

	// Ожидание завершения всех серверов
	wg.Wait()
	l.LogInfo("Все серверы остановлены. Приложение завершено.")
}
