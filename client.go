package labstack

import (
	"github.com/go-resty/resty/v2"
)

type (
	Client struct {
		key   string
		resty *resty.Client
	}

	Error struct {
		StatusCode int    `json:"status_code"`
		Code       int    `json:"code"`
		Message    string `json:"message"`
	}
)

const (
	url = "https://api.labstack.com"
)

func NewClient(key string) *Client {
	return &Client{
		key:   key,
		resty: resty.New().SetHostURL(url).SetAuthToken(key),
	}
}

func (c *Client) Currency() *CurrencyService {
	return &CurrencyService{
		resty: resty.New().SetHostURL("https://currency.labstack.com/api/v1").SetAuthToken(c.key),
	}
}

func (c *Client) Domain() *DomainService {
	return &DomainService{
		resty: resty.New().SetHostURL("https://domain.labstack.com/api/v1").SetAuthToken(c.key),
	}
}

func (c *Client) Email() *EmailService {
	return &EmailService{
		resty: resty.New().SetHostURL("https://email.labstack.com/api/v1").SetAuthToken(c.key),
	}
}

func (c *Client) IP() *IPService {
	return &IPService{
		resty: resty.New().SetHostURL("https://ip.labstack.com/api/v1").SetAuthToken(c.key),
	}
}

func (c *Client) Webpage() *WebpageService {
	return &WebpageService{
		resty: resty.New().SetHostURL("https://webpage.labstack.com/api/v1").SetAuthToken(c.key),
	}
}

func isError(status int) bool {
	return status < 200 || status >= 300
}

func (e *Error) Error() string {
	return e.Message
}
