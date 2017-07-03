package labstack

import (
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

const (
	apiURL = "https://api.labstack.com"
)

// NewClient creates a new client for the LabStack API.
func NewClient(apiKey string) *Client {
	return &Client{
		sling:  sling.New().Base(apiURL).Add("Authorization", "Bearer "+apiKey),
		logger: log.New("labstack"),
	}
}
