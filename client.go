package labstack

import (
	"time"

	"github.com/dghubble/sling"
	glog "github.com/labstack/gommon/log"
)

type (
	Client struct {
		accountID string
		apiKey    string
		sling     *sling.Sling
		logger    *glog.Logger
	}

	Fields map[string]interface{}

	SearchParameters struct {
		Query       string   `json:"query"`
		QueryString string   `json:"query_string"`
		Since       string   `json:"since"`
		Sort        []string `json:"sort"`
		Size        int      `json:"size"`
		From        int      `json:"from"`
	}
)

const (
	apiURL = "https://api.labstack.com"
)

// NewClient creates a new client for the LabStack API.
func NewClient(accountID, apiKey string) *Client {
	return &Client{
		accountID: accountID,
		apiKey:    apiKey,
		sling:     sling.New().Base(apiURL).Add("Authorization", "Bearer "+apiKey),
		logger:    glog.New("labstack"),
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

// Log returns the log service.
func (c *Client) Log() (log *Log) {
	log = &Log{
		sling:            c.sling.Path("/log"),
		logger:           c.logger,
		Level:            LevelInfo,
		Fields:           Fields{},
		BatchSize:        60,
		DispatchInterval: 60,
	}
	log.resetEntries()
	return
}

// Store returns the store service.
func (c *Client) Store() *Store {
	return &Store{
		sling:  c.sling,
		logger: c.logger,
	}
}

func (f Fields) Add(key string, value interface{}) Fields {
	f[key] = value
	return f
}

func (f Fields) Get(key string) interface{} {
	return f[key]
}
