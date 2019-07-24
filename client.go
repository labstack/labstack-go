package labstack

import (
	"github.com/go-resty/resty/v2"
)

type (
	Client struct {
		resty *resty.Client
	}

	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

const (
	url = "https://api.labstack.com"
)

func New(key string) *Client {
	return &Client{
		resty: resty.New().SetHostURL(url).SetAuthToken(key),
	}
}

func IsError(status int) bool {
	return status < 200 || status >= 300
}

func (e *Error) Error() string {
	return e.Message
}
