package mocklog

import (
	"fmt"
	"strings"
	"sync"
)

type MockLogger struct {
	mu   sync.Mutex
	logs []string
}

func (m *MockLogger) LogInfo(message string, args ...interface{}) {
	m.log("[INFO] " + fmt.Sprintf(message, args...))
}

func (m *MockLogger) LogError(message string, args ...interface{}) {
	m.log("[ERROR] " + fmt.Sprintf(message, args...))
}

func (m *MockLogger) LogFatal(message string, args ...interface{}) {
	m.log("[FATAL] " + fmt.Sprintf(message, args...))
	panic("fatal error logged")
}

func (m *MockLogger) LogWarn(message string, args ...interface{}) {
	m.log("[WARNING] " + fmt.Sprintf(message, args...))
}

func (m *MockLogger) Close() error {
	return nil
}

func (m *MockLogger) Sync() error {
	return nil
}

func (m *MockLogger) log(msg string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.logs = append(m.logs, msg)
}

func (m *MockLogger) Contains(substr string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, log := range m.logs {
		if strings.Contains(log, substr) {

			return true
		}
	}
	return false
}

func NewMockLogger() *MockLogger {
	return &MockLogger{}
}

func (m *MockLogger) ClearLogs() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.logs = []string{}
}

func (m *MockLogger) GetLogs() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]string(nil), m.logs...)
}
