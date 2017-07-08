package labstack

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/dghubble/sling"
	glog "github.com/labstack/gommon/log"
)

type (
	// Log defines the LabStack log service.
	Log struct {
		sling            *sling.Sling
		logs             []*LogEntry
		timer            <-chan time.Time
		mutex            sync.RWMutex
		logger           *glog.Logger
		AppID            string
		AppName          string
		Tags             []string
		Level            string
		BatchSize        int
		DispatchInterval int
	}

	// LogEntry defines a log entry.
	LogEntry struct {
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

func (l *Log) resetLogs() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logs = make([]*LogEntry, 0, l.BatchSize)
}

func (l *Log) appendLog(lm *LogEntry) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logs = append(l.logs, lm)
}

func (l *Log) listLogs() []*LogEntry {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	logs := make([]*LogEntry, len(l.logs))
	for i, log := range l.logs {
		logs[i] = log
	}
	return logs
}

func (l *Log) logsLength() int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return len(l.logs)
}

func (l *Log) dispatch() (err error) {
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
func (l *Log) Debug(format string, args ...interface{}) {
	l.Log(DEBUG, format, args...)
}

// Info logs an informational message.
func (l *Log) Info(format string, args ...interface{}) {
	l.Log(INFO, format, args...)
}

// Warn logs a warning message.
func (l *Log) Warn(format string, args ...interface{}) {
	l.Log(WARN, format, args...)
}

// Error logs an error message.
func (l *Log) Error(format string, args ...interface{}) {
	l.Log(ERROR, format, args...)
}

// Log logs a message with log level.
func (l *Log) Log(level, format string, args ...interface{}) {
	if levels[level] < levels[l.Level] {
		return
	}

	if l.timer == nil {
		l.timer = time.Tick(time.Duration(l.DispatchInterval) * time.Second)
		go func() {
			for range l.timer {
				if err := l.dispatch(); err != nil {
					l.logger.Error(err)
				}
			}
		}()
	}

	message := fmt.Sprintf(format, args...)
	lm := &LogEntry{
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
