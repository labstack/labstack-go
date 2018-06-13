package bolt

import (
	"os"

	"github.com/go-errors/errors"
	"github.com/go-resty/resty"
	"github.com/labstack/gommon/log"
)

type (
	Bolt struct {
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
		Version string `json:"version"`
	}

	Device struct {
		Hostname string `json:"hostname"`
	}
)

var (
	b *Bolt
)

// Severity
const (
	SeverityInfo  = Severity("info")
	SeverityWarn  = Severity("warn")
	SeverityError = Severity("error")
)

func Register(apiKey string) {
	RegisterWithOptions(apiKey, Options{})
}

func RegisterWithOptions(apiKey string, options Options) {
	b = &Bolt{
		client: resty.New().
			SetHostURL("https://api.labstack.com").
			SetAuthToken(apiKey).
			SetHeader("User-Agent", "labstack/bolt"),
		logger: log.New("bolt"),
	}
	b.Options = options

	// Defaults
	if b.Device == nil {
		b.Device = new(Device)
	}
	if b.Device.Hostname == "" {
		b.Device.Hostname, _ = os.Hostname()
	}
}

func (b *Bolt) dispatch(err *errors.Error, data *Data) {
	// fmt.Println(err.ErrorStack())
	event := newEvent(err, data)
	r, e := b.client.R().
		SetBody(event).
		// SetError(err).
		Post("/bolt")
	if e != nil {
		b.logger.Error(err)
		return
	}
	if r.StatusCode() < 200 || r.StatusCode() >= 300 {
		b.logger.Errorf("Failed to send error report: %s", r.Body())
	}
}

func Report(err error) {
	ReportWithData(err, new(Data))
}

func ReportWithData(err error, data *Data) {
	if ee, ok := err.(*errors.Error); ok {
		b.dispatch(ee, data)
	} else {
		b.dispatch(errors.Wrap(err, 1), data)
	}
}

func AutoReport() {
	if err := recover(); err != nil {
		ReportWithData(errors.Wrap(err, 1), new(Data))
		panic(err)
	}
}

func AutoReportWithData(data *Data) {
	if err := recover(); err != nil {
		ReportWithData(errors.Wrap(err, 1), data)
		panic(err)
	}
}

func Recover() {
	if err := recover(); err != nil {
		ReportWithData(errors.Wrap(err, 1), new(Data))
	}
}

func RecoverWithData(data *Data) {
	if err := recover(); err != nil {
		ReportWithData(errors.Wrap(err, 1), data)
	}
}
