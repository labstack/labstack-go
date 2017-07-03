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
		Module           string
		Level            string
		BatchSize        int
		DispatchInterval int
	}

	// Log defines a log message.
	Log struct {
		ID      string `json:"id,omitempty"`
		Time    string `json:"time"`
		Module  string `json:"module"`
		Level   string `json:"level"`
		Message string `json:"message"`
	}
)

// Log levels
const (
	DEBUG = "debug"
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
)

var levels = map[string]int{
	"debug": 1,
	"info":  2,
	"warn":  3,
	"error": 4,
}

func (l *Logging) resetLogs() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logs = make([]*Log, 0, l.BatchSize)
}

func (l *Logging) appendLog(log *Log) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logs = append(l.logs, log)
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

// Logging returns the logging service.
func (c *Client) Logging() (logging *Logging) {
	logging = &Logging{
		sling:            c.sling.Path("/logging"),
		logger:           c.logger,
		Module:           "*",
		Level:            INFO,
		BatchSize:        60,
		DispatchInterval: 60,
	}
	logging.resetLogs()
	go func() {
		d := time.Duration(logging.DispatchInterval) * time.Second
		for range time.Tick(d) {
			logging.dispatch()
		}
	}()
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
	log := &Log{
		Time:    time.Now().Format(rfc3339Milli),
		Module:  l.Module,
		Level:   level,
		Message: message,
	}
	l.appendLog(log)

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
