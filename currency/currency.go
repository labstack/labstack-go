package currency

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go/common"
	"strconv"
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
		Time   string  `json:"time"`
		Amount float64 `json:"amount"`
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

func (c *Client) Convert(req *ConvertRequest) (*ConvertResponse, *common.Error) {
	res := new(ConvertResponse)
	err := new(common.Error)
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
		return nil, &common.Error{
			Message: e.Error(),
		}
	}
	if common.IsError(r) {
		return nil, err
	}
	return res, nil
}

func (c *Client) List(req *ListRequest) (*ListResponse, *common.Error) {
	res := new(ListResponse)
	err := new(common.Error)
	r, e := c.resty.R().
		SetResult(res).
		SetError(err).
		Get("/list")
	if e != nil {
		return nil, &common.Error{
			Message: e.Error(),
		}
	}
	if common.IsError(r) {
		return nil, err
	}
	return res, nil
}
