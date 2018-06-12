package bolt

import (
	"path/filepath"

	"github.com/go-errors/errors"
	labstack "github.com/labstack/labstack-go"
)

type (
	Event struct {
		*Data
		App       *App       `json:"app"`
		Device    *Device    `json:"device"`
		Exception *Exception `json:"exception"`
	}

	Data struct {
		Severity Severity     `json:"severity"`
		User     *User        `json:"user"`
		Request  *Request     `json:"request"`
		Tags     []string     `json:"tags"`
		Extra    labstack.Map `json:"extra"`
	}

	Severity string

	User struct {
		ID string `json:"id"`
	}

	Request struct {
	}

	Exception struct {
		Class      string        `json:"class"`
		Message    string        `json:"message"`
		StackTrace []*StackFrame `json:"stack_trace"`
	}

	StackFrame struct {
		File     string         `json:"file"`
		Line     int            `json:"line"`
		Column   int            `json:"column"`
		Function string         `json:"function"`
		Code     map[int]string `json:"code"`
	}
)

func newEvent(err *errors.Error, data *Data) (e *Event) {
	e = &Event{
		Data:   data,
		App:    b.App,
		Device: b.Device,
		Exception: &Exception{
			Class:      err.TypeName(),
			Message:    err.Error(),
			StackTrace: make([]*StackFrame, len(err.StackFrames())),
		},
	}

	// Defaults
	if e.Severity == "" {
		e.Severity = SeverityWarn
	}

	for i, f := range err.StackFrames() {
		e.Exception.StackTrace[i] = &StackFrame{
			File:     filepath.Base(f.File),
			Line:     f.LineNumber,
			Function: f.Name,
		}
	}

	return
}
