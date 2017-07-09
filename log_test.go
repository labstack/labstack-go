package labstack

import (
	"testing"
)

func TestLog(t *testing.T) {
	// // dispatch
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	logs := []*logEntry{}
	// 	if err := json.NewDecoder(r.Body).Decode(&logs); err == nil {
	// 		if assert.Len(t, logs, 1) {
	// 			assert.Equal(t, LogLevelInfo, logs[0].Level)
	// 			assert.Equal(t, LogLevelInfo, logs[0].Message)
	// 			w.WriteHeader(http.StatusNoContent)
	// 			return
	// 		}
	// 	}
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }
	// ts := httptest.NewServer(http.HandlerFunc(handler))
	// defer ts.Close()
	// apiURL = ts.URL
	// l := NewClient("").Log()
	// l.appendLog(&logEntry{
	// 	Level:   levels[LogLevelInfo],
	// 	Message: levels[LogLevelInfo],
	// })
	// assert.NoError(t, l.dispatch())

	// // Log
	// for _, level := range []LogLevel{LogLevelDebug, LogLevelInfo, LogLevelWarn, LogLevelError, LogLevelFatal} {
	// 	l := NewClient("").Log()
	// 	l.Level = LogLevelDebug
	// 	switch level {
	// 	case LogLevelDebug:
	// 		l.Debug("%s", LogLevelDebug)
	// 	case LogLevelInfo:
	// 		l.Info("%s", LogLevelInfo)
	// 	case LogLevelWarn:
	// 		l.Warn("%s", LogLevelWarn)
	// 	case LogLevelError:
	// 		l.Error("%s", LogLevelError)
	// 	case LogLevelFatal:
	// 		l.Error("%s", LogLevelFatal)
	// 	}
	// 	if assert.Equal(t, 1, l.logsLength()) {
	// 		log := l.logs[0]
	// 		assert.Equal(t, levels[level], log.Level)
	// 		assert.Equal(t, levels[level], log.Message)
	// 	}
	// }

	// // Level
	// l = NewClient("").Log()
	// l.Level = LogLevelError
	// l.Debug("debug")
	// l.Info("info")
	// l.Warn("warn")
	// assert.Equal(t, 0, l.logsLength())
	// l.Level = LogLevelInfo
	// l.Info("info")
	// l.Warn("warn")
	// assert.Equal(t, 2, l.logsLength())
}
