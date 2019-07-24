package labstack

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go/currency"
)

type (
	Client struct {
		key           string
		apiResty      *resty.Client
		currencyResty *resty.Client
		domainResty   *resty.Client
		emailResty    *resty.Client
		ipResty       *resty.Client
		webpageResty  *resty.Client
	}

)

func New(key string) *Client {
	return &Client{
		key:           key,
		apiResty:      resty.New().SetHostURL("https://api.labstack.com").SetAuthToken(key),
		currencyResty: resty.New().SetHostURL("https://currency.labstack.com/api/v1").SetAuthToken(key),
		domainResty:   resty.New().SetHostURL("https://domain.labstack.com/api/v1").SetAuthToken(key),
		emailResty:    resty.New().SetHostURL("https://email.labstack.com/api/v1").SetAuthToken(key),
		ipResty:       resty.New().SetHostURL("https://ip.labstack.com/api/v1").SetAuthToken(key),
		webpageResty:  resty.New().SetHostURL("https://webpage.labstack.com/api/v1").SetAuthToken(key),
	}
}

func (c *Client) Currency() *currency.Client {
	return currency.New(c.key)
}

