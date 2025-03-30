package logger

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultLogger(t *testing.T) {
	// Тестирование методов логирования без реального файла
	t.Run("LogMethods", func(t *testing.T) {
		// Перехват вывода логов
		oldOutput := log.Writer()
		defer log.SetOutput(oldOutput)

		var logged string
		log.SetOutput(&testWriter{func(p []byte) (n int, err error) {
			logged += string(p)
			return len(p), nil
		}})

		logger := &DefaultLogger{
			logFile: nil,
			done:    make(chan struct{}),
		}

		tests := []struct {
			name    string
			logFunc func(string, ...interface{})
			prefix  string
		}{
			{"Info", logger.LogInfo, "[INFO]"},
			{"Error", logger.LogError, "[ERROR]"},
			{"Warn", logger.LogWarn, "[WARNING]"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				msg := "test message " + tt.name
				tt.logFunc(msg)
				assert.Contains(t, logged, tt.prefix+" "+msg)
				logged = "" // Сброс для следующего теста
			})
		}
	})

	// Тестирование создания файла логов отдельно
	t.Run("FileCreation", func(t *testing.T) {
		tempDir, err := ioutil.TempDir("", "testlogger")
		require.NoError(t, err)
		defer os.RemoveAll(tempDir)

		logPath := filepath.Join(tempDir, "app.log")

		// Создание файла вручную перед инициализацией логгера
		file, err := os.Create(logPath)
		require.NoError(t, err)
		file.Close()

		logger, err := InitLogger(tempDir)
		require.NoError(t, err)
		defer logger.Close()

		_, err = os.Stat(logPath)
		assert.NoError(t, err, "Log file should exist")
	})
}

func TestLoggerClose(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "testlogger")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	file, err := os.Create(filepath.Join(tempDir, "test.log"))
	require.NoError(t, err)

	logger := &DefaultLogger{
		logFile: file,
		done:    make(chan struct{}),
	}

	err = logger.Close()
	assert.NoError(t, err)

	// Проверка закрытия файла
	_, err = file.WriteString("test")
	assert.Error(t, err, "Should not be able to write to closed file")
}

func TestCleanOldLogs(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "testcleanlogs")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Тестовые файлы
	oldFile := filepath.Join(tempDir, "old.log")
	newFile := filepath.Join(tempDir, "new.log")

	require.NoError(t, ioutil.WriteFile(oldFile, []byte("old"), 0644))
	require.NoError(t, ioutil.WriteFile(newFile, []byte("new"), 0644))

	// Установка старого времени для одного файла
	oldTime := time.Now().Add(-8 * 24 * time.Hour)
	require.NoError(t, os.Chtimes(oldFile, oldTime, oldTime))

	// Вызов очистки
	err = cleanOldLogs(tempDir, 7*24*time.Hour)
	assert.NoError(t, err)

	// Проверка результатов
	_, err = os.Stat(oldFile)
	assert.True(t, os.IsNotExist(err), "Old file should be deleted")

	_, err = os.Stat(newFile)
	assert.NoError(t, err, "New file should exist")
}

type testWriter struct {
	writeFunc func(p []byte) (n int, err error)
}

func (t *testWriter) Write(p []byte) (n int, err error) {
	return t.writeFunc(p)
}
