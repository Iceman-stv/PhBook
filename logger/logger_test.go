package logger

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// createTestDir создает временную директорию для тестов и регистрирует её удаление
// с несколькими попытками в случае ошибки. Возвращает путь к созданной директории.
func createTestDir(t *testing.T, prefix string) string {
	tempDir, err := ioutil.TempDir("", prefix)
	if err != nil {

		t.Fatalf("Не удалось создать временную директорию: %v", err)
	}

	// Регистрация очистки директории после завершения теста
	t.Cleanup(func() {
		maxAttempts := 3  // Максимальное количество попыток удаления
		var lastErr error // Для хранения последней ошибки
		success := false  // Флаг успешного удаления

		// Вывод содержимого директории при неудачных тестах
		if t.Failed() {

			files, _ := ioutil.ReadDir(tempDir)
			t.Logf("Содержимое директории перед удалением (%s): %v", tempDir, files)
		}

		// Попытка удалить директорию несколько раз
		for attempt := 1; attempt <= maxAttempts; attempt++ {
			err := os.RemoveAll(tempDir)
			if err == nil {

				success = true
				break
			}
			lastErr = err
			time.Sleep(100 * time.Millisecond) // Задержка между попытками
		}

		if !success {

			t.Logf("Не удалось удалить директорию %s после %d попыток: %v",
				tempDir, maxAttempts, lastErr)
		}
	})

	return tempDir
}

// verifyDirEmpty проверяет, что директория пуста, и логирует содержимое если это не так
func verifyDirEmpty(t *testing.T, dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {

		t.Logf("Ошибка при проверке директории %s: %v", dir, err)
		return
	}

	if len(files) > 0 {

		t.Logf("Директория %s не пуста после теста. Оставшиеся файлы:", dir)
		for _, f := range files {
			t.Logf(" - %s (size: %d, mod: %v)",
				f.Name(), f.Size(), f.ModTime().Format(time.RFC3339))
		}
	}
}

func TestLogMethods(t *testing.T) {
	tempDir := createTestDir(t, "test_logs")

	// Инициализация логгера
	logger, err := InitLogger(tempDir)
	if err != nil {

		t.Fatalf("Ошибка инициализации логгера: %v", err)
	}

	// Гарантированное закрытие логгера после теста
	t.Cleanup(func() {
		if err := logger.Close(); err != nil {

			t.Errorf("Ошибка при закрытии логгера: %v", err)
		}

		// Дополнительная проверка что файл действительно закрыт
		if dl, ok := logger.(*DefaultLogger); ok && dl.logFile != nil {

			t.Error("Файл лога не был закрыт полностью")
		}
	})

	// Тест методов логирования
	logger.LogInfo("Тестовое информационное сообщение %d", 1)
	logger.LogWarn("Тестовое предупреждение %d", 2)
	logger.LogError("Тестовая ошибка %d", 3)

	// Принудительная запись логов на диск
	if dl, ok := logger.(*DefaultLogger); ok {

		if err := dl.Sync(); err != nil {

			t.Errorf("Ошибка синхронизации логов: %v", err)
		}
	}

	// Проверка записи в файл
	logFilePath := filepath.Join(tempDir, "app.log")
	content, err := ioutil.ReadFile(logFilePath)
	if err != nil {

		t.Fatalf("Ошибка чтения лог-файла: %v", err)
	}

	// Проверки содержимого логов
	testCases := []struct {
		level    string
		message  string
		expected string
	}{
		{"INFO", "Тестовое информационное сообщение 1", "[INFO] Тестовое информационное сообщение 1"},
		{"WARNING", "Тестовое предупреждение 2", "[WARNING] Тестовое предупреждение 2"},
		{"ERROR", "Тестовая ошибка 3", "[ERROR] Тестовая ошибка 3"},
	}

	logContent := string(content)
	for _, tc := range testCases {
		if !strings.Contains(logContent, tc.expected) {

			t.Errorf("Сообщение уровня %s не найдено в логе: %q", tc.level, tc.message)
		}
	}

	// Тест LogFatal в отдельном процессе
	if os.Getenv("TEST_FATAL") == "1" {

		logger.LogFatal("Тестовая фатальная ошибка %d", 4)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestLogMethods")
	cmd.Env = append(os.Environ(), "TEST_FATAL=1")
	err = cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {

		// Ожидаемое поведение - LogFatal должен завершить процесс с кодом 1
	} else {
		t.Errorf("LogFatal не завершил процесс с кодом 1: %v", err)
	}

	// Проверка что директория не содержит лишних файлов
	t.Cleanup(func() { verifyDirEmpty(t, tempDir) })
}

func TestCleanOldLogs(t *testing.T) {
	tempDir := createTestDir(t, "test_cleanup")

	// Тестовые файлы логов
	oldFile := filepath.Join(tempDir, "old.log")
	newFile := filepath.Join(tempDir, "new.log")

	// Создание старог файла (8 дней назад)
	if err := ioutil.WriteFile(oldFile, []byte("old logs"), 0644); err != nil {

		t.Fatalf("Ошибка создания старого лог-файла: %v", err)
	}
	oldTime := time.Now().Add(-8 * 24 * time.Hour)
	if err := os.Chtimes(oldFile, oldTime, oldTime); err != nil {

		t.Fatalf("Ошибка изменения времени файла: %v", err)
	}

	// Создание нового файла
	if err := ioutil.WriteFile(newFile, []byte("new logs"), 0644); err != nil {

		t.Fatalf("Ошибка создания нового лог-файла: %v", err)
	}

	// Тест очистки логов (старше 7 дней)
	if err := cleanOldLogs(tempDir, 7*24*time.Hour); err != nil {

		t.Fatalf("Ошибка очистки логов: %v", err)
	}

	// Проверка что старый файл удален
	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {

		t.Error("Старый лог-файл не был удален")
	}

	// Проверка что новый файл остался
	if _, err := os.Stat(newFile); err != nil {

		t.Error("Новый лог-файл был удален")
	}
}

func TestLogClose(t *testing.T) {
	tempDir := createTestDir(t, "test_close")

	// Инициализация логгера
	logger, err := InitLogger(tempDir)
	if err != nil {

		t.Fatalf("Ошибка инициализации логгера: %v", err)
	}

	// Закрывание логгера
	if err := logger.Close(); err != nil {

		t.Errorf("Ошибка при закрытии логгера: %v", err)
	}

	// Попытка записать в закрытый файл
	if dl, ok := logger.(*DefaultLogger); ok && dl.logFile != nil {

		_, err = dl.logFile.WriteString("test")
		if err == nil {

			t.Error("Запись в закрытый файл не вернула ошибку")
		}
	}
}

func TestLogOutput(t *testing.T) {
	tempDir := createTestDir(t, "test_output")

	// Инициализация логгера
	logger, err := InitLogger(tempDir)
	if err != nil {

		t.Fatalf("Ошибка инициализации логгера: %v", err)
	}
	t.Cleanup(func() {
		if err := logger.Close(); err != nil {

			t.Logf("Ошибка при закрытии логгера: %v", err)
		}
	})

	// Перехват вывода стандартного log
	log.SetOutput(logger.(*DefaultLogger).logFile)

	// Тестовое сообщение
	testMessage := "Тестовое сообщение в лог"
	log.Print(testMessage)

	// Синхронизация записи
	if err := logger.(*DefaultLogger).Sync(); err != nil {

		t.Fatalf("Ошибка синхронизации логов: %v", err)
	}

	// Проверка записи в файл
	logFilePath := filepath.Join(tempDir, "app.log")
	content, err := ioutil.ReadFile(logFilePath)
	if err != nil {
		t.Fatalf("Ошибка чтения лог-файла: %v", err)
	}

	if !strings.Contains(string(content), testMessage) {

		t.Error("Вывод логов не перенаправлен в файл")
	}
}
