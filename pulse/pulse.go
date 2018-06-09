package pulse

import (
	"fmt"
	"os"

	"github.com/go-errors/errors"
	"github.com/go-resty/resty"
	"github.com/labstack/gommon/log"
)

type (
	Pulse struct {
		Options
		// mutex  sync.RWMutex
		client *resty.Client
		logger *log.Logger
	}

	Options struct {
		App    *App
		Device *Device
	}

	App struct {
		Version string
	}

	Device struct {
		Hostname string
	}
)

var (
	p *Pulse
)

// Severity
var (
	SeverityInfo  = severity("info")
	SeverityWarn  = severity("warn")
	SeverityError = severity("error")
)

func Register(apiKey string) {
	RegisterWithOptions(apiKey, Options{})
}

func RegisterWithOptions(apiKey string, options Options) {
	p = &Pulse{
		client: resty.New().
			SetHostURL("https://api.labstack.com").
			SetAuthToken(apiKey).
			SetHeader("User-Agent", "labstack/pulse"),
		logger: log.New("pulse"),
	}
	p.Options = options

	// Defaults
	if p.Device == nil {
		p.Device = new(Device)
	}
	if p.Device.Hostname == "" {
		p.Device.Hostname, _ = os.Hostname()
	}
}

func (p *Pulse) dispatch(err *errors.Error, ctx *Context) {
	fmt.Println(err.ErrorStack())
	// event := newEvent(err, ctx)
	// b, _ := json.MarshalIndent(event, "", " ")
	// fmt.Printf("%s", b)
	// for _, f := range e.StackFrames() {
	// 	fmt.Println(f.File, f.LineNumber, f.Name)
	// }
	// res, err := p.client.R().
	// 	SetBody(event).
	// 	// SetError(err).
	// 	Post("/pulse")
	// if err != nil {
	// 	p.logger.Error(err)
	// 	return
	// }
	// if res.StatusCode() < 200 || res.StatusCode() >= 300 {
	// 	p.logger.Error(res.Body())
	// }
}

func Report(err error) {
	ReportWithContext(err, new(Context))
}

func ReportWithContext(err error, ctx *Context) {
	if ee, ok := err.(*errors.Error); ok {
		p.dispatch(ee, ctx)
	} else {
		p.dispatch(errors.Wrap(err, 1), ctx)
	}
}

func AutoReport() {
	if err := recover(); err != nil {
		ReportWithContext(errors.Wrap(err, 1), new(Context))
		panic(err)
	}
}

func AutoReportWithContext(ctx *Context) {
	if err := recover(); err != nil {
		ReportWithContext(errors.Wrap(err, 1), ctx)
		panic(err)
	}
}

func Recover() {
	if err := recover(); err != nil {
		ReportWithContext(errors.Wrap(err, 1), new(Context))
	}
}

func RecoverWithContext(ctx *Context) {
	if err := recover(); err != nil {
		ReportWithContext(errors.Wrap(err, 1), ctx)
	}
}
