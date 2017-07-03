package labstack

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	// dispatch
	handler := func(w http.ResponseWriter, r *http.Request) {
		logs := []*Log{}
		if err := json.NewDecoder(r.Body).Decode(&logs); err == nil {
			if assert.Len(t, logs, 1) {
				assert.Equal(t, INFO, logs[0].Level)
				assert.Equal(t, INFO, logs[0].Message)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()
	apiURL = ts.URL
	l := NewClient("").Logging()
	l.appendLog(&Log{
		Level:   INFO,
		Message: INFO,
	})
	assert.NoError(t, l.dispatch())

	// Log
	for _, level := range []string{DEBUG, INFO, WARN, ERROR} {
		l := NewClient("").Logging()
		l.Level = DEBUG
		l.BatchSize = 1
		switch level {
		case DEBUG:
			l.Debug("%s", DEBUG)
		case INFO:
			l.Info("%s", INFO)
		case WARN:
			l.Warn("%s", WARN)
		case ERROR:
			l.Error("%s", ERROR)
		}
		if assert.Equal(t, 1, l.logsLength()) {
			log := l.logs[0]
			assert.Equal(t, level, log.Level)
			assert.Equal(t, level, log.Message)
		}
	}

	// Level
	l = NewClient("").Logging()
	l.Level = ERROR
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	assert.Equal(t, 0, l.logsLength())
	l.Level = INFO
	l.Info("info")
	l.Warn("warn")
	assert.Equal(t, 2, l.logsLength())
}
