package pulse

import (
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
		Device *Device
	}

	User struct {
	}

	Device struct {
	}

	Data struct {
	}

	Context struct {
		user *User
		data *Data
	}

	// APIError struct {
	// 	Code    int    `json:"code"`
	// 	Message string `json:"message"`
	// }
)

var (
	global *Pulse
)

func Register(apiKey string) {
	RegisterWithOptions(apiKey, Options{})
}

func RegisterWithOptions(apiKey string, options Options) {
	global = &Pulse{
		client: resty.New().
			SetHostURL("https://api.labstack.com").
			SetAuthToken(apiKey).
			SetHeader("User-Agent", "labstack/pulse"),
		logger: log.New("pulse"),
	}
	global.Options = options

	// Defaults
}

func (p *Pulse) Report() *Context {
	return new(Context)
}

func (p *Pulse) AutoReport() {
}

func (p *Pulse) Recover() {
}

func (c *Context) SetUser(u *User) *Context {
	c.user = u
	return c
}

func (c *Context) SetData(d *Data) *Context {
	c.data = d
	return c
}

// // Dispatch dispatches the requests batch.
// func (c *Cube) Dispatch() {
// 	if len(c.requests) == 0 {
// 		return
// 	}

// 	// err := new(APIError)
// 	res, err := c.client.R().
// 		SetBody(c.readRequests()).
// 		// SetError(err).
// 		Post("/cube")
// 	if err != nil {
// 		c.logger.Error(err)
// 		return
// 	}
// 	if res.StatusCode() < 200 || res.StatusCode() >= 300 {
// 		c.logger.Error(res.Body())
// 	}

// 	c.resetRequests()
// }
