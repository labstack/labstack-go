package currency

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go"
	"strconv"
	"time"
)

type (
	Client struct {
		resty *resty.Client
	}

	Currency struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Symbol string `json:"symbol"`
	}

	ConvertRequest struct {
		Amount float64
		From   string
		To     string
	}

	ConvertResponse struct {
		Time   time.Time `json:"time"`
		Amount float64   `json:"amount"`
	}

	ListRequest struct {
	}

	ListResponse struct {
		Currencies []*Currency `json:"currencies"`
	}
)

const (
	url = "https://currency.labstack.com/api/v1"
)

func New(key string) *Client {
	return &Client{
		resty: resty.New().SetHostURL(url).SetAuthToken(key),
	}
}

func (c *Client) Convert(req *ConvertRequest) (*ConvertResponse, error) {
	res := new(ConvertResponse)
	err := new(labstack.Error)
	r, e := c.resty.R().
		SetPathParams(map[string]string{
			"amount": strconv.FormatFloat(req.Amount, 'f', -1, 64),
			"from":   req.From,
			"to":     req.To,
		}).
		SetResult(res).
		SetError(err).
		Get("/convert/{amount}/{from}/{to}")
	if e != nil {
		return nil, &labstack.Error{
			Message: e.Error(),
		}
	}
	if labstack.IsError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}

func (c *Client) List(req *ListRequest) (*ListResponse, error) {
	res := new(ListResponse)
	err := new(labstack.Error)
	r, e := c.resty.R().
		SetResult(res).
		SetError(err).
		Get("/list")
	if e != nil {
		return nil, &labstack.Error{
			Message: e.Error(),
		}
	}
	if labstack.IsError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}
