package labstack

import "time"

type (
	CurrencyConvertRequest struct {
		Base string `json:"base"`
	}

	CurrencyConvertResponse struct {
		Rates     map[string]float64 `json:"rates"`
		UpdatedAt time.Time          `json:"updated_at"`
	}
)

func (c *Client) CurrencyConvert(req *CurrencyConvertRequest) (*CurrencyConvertResponse, *APIError) {
	res := new(CurrencyConvertResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/currency/exchange")
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
