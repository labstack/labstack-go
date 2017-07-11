package labstack

import (
	"time"

	"github.com/dghubble/sling"
	glog "github.com/labstack/gommon/log"
)

type (
	Client struct {
		sling   *sling.Sling
		logger  *glog.Logger
		AppID   string
		AppName string
	}
)

const (
	apiURL = "https://api.labstack.com"
)

// NewClient creates a new client for the LabStack API.
func NewClient(apiKey string) *Client {
	return &Client{
		sling:  sling.New().Base(apiURL).Add("Authorization", "Bearer "+apiKey),
		logger: glog.New("labstack"),
	}
}

// Cube returns the cube service.
func (c *Client) Cube() (cube *Cube) {
	cube = &Cube{
		sling:            c.sling.Path("/cube"),
		logger:           c.logger,
		AppID:            c.AppID,
		AppName:          c.AppName,
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

// Log returns the log service.
func (c *Client) Log() (log *Log) {
	log = &Log{
		sling:            c.sling.Path("/log"),
		logger:           c.logger,
		AppID:            c.AppID,
		AppName:          c.AppName,
		Level:            LogLevelInfo,
		BatchSize:        60,
		DispatchInterval: 60,
	}
	log.resetLogs()
	return
}

// Store returns the store service.
func (c *Client) Store() *Store {
	return &Store{
		sling:  c.sling,
		logger: c.logger,
	}
}
