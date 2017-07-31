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
		entries          []Fields
		timer            <-chan time.Time
		mutex            sync.RWMutex
		logger           *glog.Logger
		Level            Level
		Fields           Fields
		BatchSize        int
		DispatchInterval int
	}

	// Level defines the log level.
	Level int

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

func (l *Log) resetEntries() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.entries = make([]Fields, 0, l.BatchSize)
}

func (l *Log) appendEntry(f Fields) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.entries = append(l.entries, f)
}

func (l *Log) listEntries() []Fields {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	entries := make([]Fields, len(l.entries))
	for i, entry := range l.entries {
		entries[i] = entry
	}
	return entries
}

func (l *Log) logsLength() int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return len(l.entries)
}

func (l *Log) dispatch() error {
	if len(l.entries) == 0 {
		return nil
	}

	defer l.resetEntries()

	le := new(LogError)
	_, err := l.sling.Post("").BodyJSON(l.listEntries()).Receive(nil, le)
	if err != nil {
		return err
	}
	if le.Code == 0 {
		return nil
	}
	return le
}

// Debug logs a message with DEBUG level.
func (l *Log) Debug(fields Fields) {
	l.Log(LevelDebug, fields)
}

// Info logs a message with INFO level.
func (l *Log) Info(fields Fields) {
	l.Log(LevelInfo, fields)
}

// Warn logs a message with WARN level.
func (l *Log) Warn(fields Fields) {
	l.Log(LevelWarn, fields)
}

// Error logs a message with ERROR level.
func (l *Log) Error(fields Fields) {
	l.Log(LevelError, fields)
}

// Fatal logs a message with FATAL level.
func (l *Log) Fatal(fields Fields) {
	l.Log(LevelFatal, fields)
}

// Log logs a message with log level.
func (l *Log) Log(level Level, fields Fields) {
	if level < l.Level {
		return
	}

	if l.timer == nil {
		l.timer = time.Tick(time.Duration(l.DispatchInterval) * time.Second)
		go func() {
			for range l.timer {
				if err := l.dispatch(); err != nil {
					err := err.(*LogError)
					fmt.Printf("log error: code=%d, message=%s", err.Code, err.Message)
				}
			}
		}()
	}

	fields.Add("time", time.Now().Format(rfc3339Milli)).
		Add("level", levels[level])
	for k, v := range l.Fields {
		fields.Add(k, v)
	}
	l.appendEntry(fields)

	// Dispatch batch
	if l.logsLength() >= l.BatchSize {
		go func() {
			if err := l.dispatch(); err != nil {
				err := err.(*LogError)
				fmt.Printf("log error: code=%d, message=%s", err.Code, err.Message)
			}
		}()
	}
}

func (e *LogError) Error() string {
	return e.Message
}
