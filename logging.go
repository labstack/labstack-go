package labstack

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/dghubble/sling"
	"github.com/labstack/gommon/log"
)

type (
	// Logging defines the LabStack logging service.
	Logging struct {
		sling            *sling.Sling
		logs             []*Log
		mutex            sync.RWMutex
		logger           *log.Logger
		AppID            string
		AppName          string
		Tags             []string
		Level            string
		BatchSize        int
		DispatchInterval int
	}

	// Log defines a log message.
	Log struct {
		// ID      string `json:"id,omitempty"`
		Time    string   `json:"time"`
		AppID   string   `json:"app_id"`
		AppName string   `json:"app_name"`
		Tags    []string `json:"tags"`
		Level   string   `json:"level"`
		Message string   `json:"message"`
	}
)

// Log levels
const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

var levels = map[string]int{
	"DEBUG": 1,
	"INFO":  2,
	"WARN":  3,
	"ERROR": 4,
}

func (l *Logging) resetLogs() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logs = make([]*Log, 0, l.BatchSize)
}

func (l *Logging) appendLog(lm *Log) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logs = append(l.logs, lm)
}

func (l *Logging) listLogs() []*Log {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	logs := make([]*Log, len(l.logs))
	for i, log := range l.logs {
		logs[i] = log
	}
	return logs
}

func (l *Logging) logsLength() int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return len(l.logs)
}

func (l *Logging) dispatch() (err error) {
	if len(l.logs) == 0 {
		return
	}
	res, err := l.sling.Post("").BodyJSON(l.listLogs()).Receive(nil, nil)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("logging: error dispatching logs, status=%d, message=%v", res.StatusCode, err)
	}
	return
}

// Debug logs a debug message.
func (l *Logging) Debug(format string, args ...interface{}) {
	l.Log(DEBUG, format, args...)
}

// Info logs an informational message.
func (l *Logging) Info(format string, args ...interface{}) {
	l.Log(INFO, format, args...)
}

// Warn logs a warning message.
func (l *Logging) Warn(format string, args ...interface{}) {
	l.Log(WARN, format, args...)
}

// Error logs an error message.
func (l *Logging) Error(format string, args ...interface{}) {
	l.Log(ERROR, format, args...)
}

// Log logs a message with log level.
func (l *Logging) Log(level, format string, args ...interface{}) {
	if levels[level] < levels[l.Level] {
		return
	}
	message := fmt.Sprintf(format, args...)
	lm := &Log{
		Time:    time.Now().Format(rfc3339Milli),
		AppID:   l.AppID,
		AppName: l.AppName,
		Tags:    l.Tags,
		Level:   level,
		Message: message,
	}
	l.appendLog(lm)

	// Dispatch batch
	if l.logsLength() >= l.BatchSize {
		go func() {
			if err := l.dispatch(); err != nil {
				l.logger.Error(err)
			}
			l.resetLogs()
		}()
	}
}
