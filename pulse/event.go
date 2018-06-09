package pulse

import "github.com/go-errors/errors"

type (
	Event struct {
		*Context
		App       *App       `json:"app" db:"app"`
		Device    *Device    `json:"device" db:"device"`
		Exception *Exception `json:"exception" db:"exception"`
	}

	Context struct {
		Severity severity `json:"severity" db:"severity"`
		User     *User    `json:"user" db:"user"`
		Request  *Request `json:"request" db:"request"`
		Data     Data     `json:"data" db:"data"`
	}

	severity string

	User struct {
		ID string `json:"id"`
	}

	Request struct {
	}

	Data map[string]interface{}

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

func newEvent(err *errors.Error, ctx *Context) (e *Event) {
	e = &Event{
		Context: ctx,
		App:     p.App,
		Device:  p.Device,
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
			File:     f.File,
			Line:     f.LineNumber,
			Function: f.Name,
		}
	}

	return
}
