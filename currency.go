package labstack

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go/currency"
	"strconv"
)

type (
	CurrencyService struct {
		resty *resty.Client
	}
)

func (c *CurrencyService) Convert(req *currency.ConvertRequest) (*currency.ConvertResponse, error) {
	res := new(currency.ConvertResponse)
	err := new(Error)
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
		return nil, &Error{
			Message: e.Error(),
		}
	}
	if isError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}

func (c *CurrencyService) List(req *currency.ListRequest) (*currency.ListResponse, error) {
	res := new(currency.ListResponse)
	err := new(Error)
	r, e := c.resty.R().
		SetResult(res).
		SetError(err).
		Get("/list")
	if e != nil {
		return nil, &Error{
			Message: e.Error(),
		}
	}
	if isError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}

func (c *CurrencyService) Rates(req *currency.RatesRequest) (*currency.RatesResponse, error) {
	res := new(currency.RatesResponse)
	err := new(Error)
	r, e := c.resty.R().
		SetResult(res).
		SetError(err).
		Get("/rates")
	if e != nil {
		return nil, &Error{
			Message: e.Error(),
		}
	}
	if isError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}
