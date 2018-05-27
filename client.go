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

	Download struct {
		ID   string `json:"id"`
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

func (c *Client) error(r *resty.Response) bool {
	return r.StatusCode() < 200 || r.StatusCode() >= 300
}

func (c *Client) Currency() *Currency {
	return &Currency{c}
}

func (c *Client) Geocode() *Geocode {
	return &Geocode{c}
}

func (c *Client) Email() *Email {
	return &Email{c}
}

func (c *Client) Download(id string, path string) (ae *APIError) {
	_, err := c.resty.R().
		SetOutput(path).
		Get(fmt.Sprintf("%s/download/%s", apiURL, id))
	if err != nil {
		ae = new(APIError)
		ae.Message = err.Error()
	}
	return
}

func (e *APIError) Error() string {
	return e.Message
}
