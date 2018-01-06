package labstack

import (
	"fmt"

	"github.com/go-resty/resty"
	"github.com/labstack/gommon/log"
)

type (
	Client struct {
		apiKey string
		resty  *resty.Client
		logger *log.Logger
	}

	Response struct {
		ID string `json:"id"`
	}

	Download struct {
		*Response
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	APIError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

const (
	apiURL = "https://api.labstack.com"
)

// NewClient creates a new client for the LabStack API.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		resty:  resty.New().SetHostURL(apiURL).SetAuthToken(apiKey),
		logger: log.New("labstack"),
	}
}

func (c *Client) Download(id string, path string) (err *APIError) {
	_, e := c.resty.R().
		SetOutput(path).
		Get(fmt.Sprintf("%s/download/%s", apiURL, id))
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}

func (e *APIError) Error() string {
	return e.Message
}
