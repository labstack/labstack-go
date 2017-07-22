package labstack

import (
	"fmt"
	"sync"
	"time"

	"github.com/dghubble/sling"
	glog "github.com/labstack/gommon/log"
)

type (
	// Log defines the LabStack log service.
	Log struct {
		sling            *sling.Sling
		logs             []*logEntry
		timer            <-chan time.Time
		mutex            sync.RWMutex
		logger           *glog.Logger
		AppID            string
		AppName          string
		Tags             []string
		Level            Level
		BatchSize        int
		DispatchInterval int
	}

	// Level defines the log level.
	Level int

	// logEntry defines a log entry.
	logEntry struct {
		Time    string   `json:"time"`
		AppID   string   `json:"app_id"`
		AppName string   `json:"app_name"`
		Tags    []string `json:"tags"`
		Level   string   `json:"level"`
		Message string   `json:"message"`
	}

	// LogError defines the log error.
	LogError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

// Log levels
const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LogLevelOff
)

var levels = map[Level]string{
	LevelDebug: "DEBUG",
	LevelInfo:  "INFO",
	LevelWarn:  "WARN",
	LevelError: "ERROR",
	LevelFatal: "FATAL",
}

func (l *Log) resetLogs() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logs = make([]*logEntry, 0, l.BatchSize)
}

func (l *Log) appendLog(lm *logEntry) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logs = append(l.logs, lm)
}

func (l *Log) listLogs() []*logEntry {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	logs := make([]*logEntry, len(l.logs))
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

func (l *Log) dispatch() error {
	if len(l.logs) == 0 {
		return nil
	}

	le := new(LogError)
	_, err := l.sling.Post("").BodyJSON(l.listLogs()).Receive(nil, le)
	if err != nil {
		return err
	}
	if le.Code == 0 {
		return nil
	}
	return le
}

// Debug logs a message with DEBUG level.
func (l *Log) Debug(format string, args ...interface{}) {
	l.Log(LevelDebug, format, args...)
}

// Info logs a message with INFO level.
func (l *Log) Info(format string, args ...interface{}) {
	l.Log(LevelInfo, format, args...)
}

// Warn logs a message with WARN level.
func (l *Log) Warn(format string, args ...interface{}) {
	l.Log(LevelWarn, format, args...)
}

// Error logs a message with ERROR level.
func (l *Log) Error(format string, args ...interface{}) {
	l.Log(LevelError, format, args...)
}

// Fatal logs a message with FATAL level.
func (l *Log) Fatal(format string, args ...interface{}) {
	l.Log(LevelFatal, format, args...)
}

// Log logs a message with log level.
func (l *Log) Log(level Level, format string, args ...interface{}) {
	if level < l.Level {
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
	lm := &logEntry{
		Time:    time.Now().Format(rfc3339Milli),
		AppID:   l.AppID,
		AppName: l.AppName,
		Tags:    l.Tags,
		Level:   levels[level],
		Message: message,
	}
	l.appendLog(lm)

	// Dispatch batch
	if l.logsLength() >= l.BatchSize {
		go func() {
			if err := l.dispatch(); err != nil {
				l.logger.Error(err)
			}
		}()
	}
}

func (e *LogError) Error() string {
	return fmt.Sprintf("log error, code=%d, message=%s", e.Code, e.Message)
}
