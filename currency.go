package labstack

import (
	"strconv"
	"time"
)

type (
	Currency struct {
		*Client
	}

	CurrencyConvertResponse struct {
		Time   time.Time `json:"time"`
		Amount float64   `json:"amount"`
	}

	CurrencyRatesResponse struct {
		Time  time.Time          `json:"time"`
		Rates map[string]float64 `json:"rates"`
	}
)

func (c *Currency) Convert(amount float64, from, to string) (*CurrencyConvertResponse, *APIError) {
	res := new(CurrencyConvertResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetQueryParams(map[string]string{
			"value": strconv.FormatFloat(amount, 'f', -1, 64),
			"from":  from,
			"to":    to,
		}).
		SetResult(res).
		SetError(err).
		Get("/currency/convert")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if c.error(r) {
		return nil, err
	}
	return res, nil
}

func (c *Currency) Rates(base string) (*CurrencyRatesResponse, *APIError) {
	res := new(CurrencyRatesResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetQueryParams(map[string]string{
			"base": base,
		}).
		SetResult(res).
		SetError(err).
		Get("/currency/rates")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if c.error(r) {
		return nil, err
	}
	return res, nil
}
