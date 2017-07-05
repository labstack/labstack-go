package labstack

import (
	"time"

	"github.com/dghubble/sling"
	"github.com/labstack/gommon/log"
)

type (
	Client struct {
		sling  *sling.Sling
		logger *log.Logger
	}

	Error struct {
	}
)

var (
	apiURL = "https://api.labstack.com"
)

// NewClient creates a new client for the LabStack API.
func NewClient(apiKey string) *Client {
	return &Client{
		sling:  sling.New().Base(apiURL).Add("Authorization", "Bearer "+apiKey),
		logger: log.New("labstack"),
	}
}

// Cube returns the cube service.
func (c *Client) Cube() (cube *Cube) {
	cube = &Cube{
		sling:            c.sling.Path("/cube"),
		logger:           c.logger,
		BatchSize:        60,
		DispatchInterval: 60,
	}
	cube.resetRequests()
	go func() {
		d := time.Duration(cube.DispatchInterval) * time.Second
		for range time.Tick(d) {
			cube.dispatch()
		}
	}()
	return
}

// Email returns the email service.
func (c *Client) Email() *Email {
	return &Email{
		sling:  c.sling.Path("/email"),
		logger: c.logger,
	}
}

// Logging returns the logging service.
func (c *Client) Logging() (logging *Logging) {
	logging = &Logging{
		sling:            c.sling.Path("/logging"),
		logger:           c.logger,
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

// Store returns the store service.
func (c *Client) Store() *Store {
	return &Store{
		sling:  c.sling,
		logger: c.logger,
	}
}
