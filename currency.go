package labstack

import (
	"github.com/labstack/labstack-go/currency"
	"strconv"
)

func (c *Client) CurrencyConvert(req *currency.ConvertRequest) (*currency.ConvertResponse, *Error) {
	res := new(currency.ConvertResponse)
	err := new(Error)
	r, e := c.currencyResty.R().
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
	if isError(r) {
		return nil, err
	}
	return res, nil
}

func (c *Client) CurrencyList(req *currency.ListRequest) (*currency.ListResponse, *Error) {
	res := new(currency.ListResponse)
	err := new(Error)
	r, e := c.currencyResty.R().
		SetResult(res).
		SetError(err).
		Get("/list")
	if e != nil {
		return nil, &Error{
			Message: e.Error(),
		}
	}
	if isError(r) {
		return nil, err
	}
	return res, nil
}
