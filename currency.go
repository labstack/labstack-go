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
		Amount    float64   `json:"amount"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	CurrencyRatesResponse struct {
		Rates     map[string]float64 `json:"rates"`
		UpdatedAt time.Time          `json:"updated_at"`
	}
)

func (c *Currency) Convert(from, to string, amount float64) (*CurrencyConvertResponse, *APIError) {
	res := new(CurrencyConvertResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetQueryParams(map[string]string{
			"from":  from,
			"to":    to,
			"value": strconv.FormatFloat(amount, 'f', -1, 64),
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
