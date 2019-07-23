package labstack

import "github.com/go-resty/resty/v2"

type (
	Client struct {
		apiResty      *resty.Client
		currencyResty *resty.Client
		domainResty   *resty.Client
		emailResty    *resty.Client
		ipResty       *resty.Client
		webpageResty  *resty.Client
	}

	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func New(key string) *Client {
	return &Client{
		apiResty:      resty.New().SetHostURL("https://api.labstack.com").SetAuthToken(key),
		currencyResty: resty.New().SetHostURL("https://currency.labstack.com/api/v1").SetAuthToken(key),
		domainResty:   resty.New().SetHostURL("https://domain.labstack.com/api/v1").SetAuthToken(key),
		emailResty:    resty.New().SetHostURL("https://email.labstack.com/api/v1").SetAuthToken(key),
		ipResty:       resty.New().SetHostURL("https://ip.labstack.com/api/v1").SetAuthToken(key),
		webpageResty:  resty.New().SetHostURL("https://webpage.labstack.com/api/v1").SetAuthToken(key),
	}
}

func (e *Error) Error() string {
	return e.Message
}
