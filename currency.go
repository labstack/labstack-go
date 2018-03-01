package labstack

import (
	"strconv"
	"time"
)

type (
	CurrencyConvertRequest struct {
		From  string
		To    string
		Value float64
	}

	CurrencyConvertResponse struct {
		Value     float64   `json:"value" db:"value"`
		UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	}

	CurrencyRatesRequest struct {
		Base string
	}

	CurrencyRatesResponse struct {
		Rates     map[string]float64 `json:"rates"`
		UpdatedAt time.Time          `json:"updated_at"`
	}
)

func (c *Client) CurrencyConvert(req *CurrencyConvertRequest) (*CurrencyConvertResponse, *APIError) {
	res := new(CurrencyConvertResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetQueryParams(map[string]string{
			"from":  req.From,
			"to":    req.To,
			"value": strconv.FormatFloat(req.Value, 'f', -1, 64),
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

func (c *Client) CurrencyRates(req *CurrencyRatesRequest) (*CurrencyRatesResponse, *APIError) {
	res := new(CurrencyRatesResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetQueryParams(map[string]string{
			"base": req.Base,
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
