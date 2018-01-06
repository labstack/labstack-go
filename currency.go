package labstack

import "time"

type (
	CurrencyExchangeRequest struct {
		Base string `json:"base"`
	}

	CurrencyExchangeResponse struct {
		*Response
		Rates     map[string]float64 `json:"rates"`
		UpdatedAt time.Time          `json:"updated_at"`
	}
)

func (c *Client) CurrencyExchange(req *CurrencyExchangeRequest) (res *CurrencyExchangeResponse, err *APIError) {
	res = new(CurrencyExchangeResponse)
	_, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/currency/exchange")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
