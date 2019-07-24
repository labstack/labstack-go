package email

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go"
)

type (
	Client struct {
		resty *resty.Client
	}

	VerifyRequest struct {
		Email string
	}

	VerifyResponse struct {
		Email    string   `json:"email"`
		Username string   `json:"username"`
		Domain   string   `json:"domain"`
		Result   string   `json:"result"`
		Flags    []string `json:"flags"`
	}
)

const (
	url = "https://email.labstack.com/api/v1"
)

func New(key string) *Client {
	return &Client{
		resty: resty.New().SetHostURL(url).SetAuthToken(key),
	}
}

func (c *Client) Verify(req *VerifyRequest) (*VerifyResponse, error) {
	res := new(VerifyResponse)
	err := new(labstack.Error)
	r, e := c.resty.R().
		SetPathParams(map[string]string{
			"email": req.Email,
		}).
		SetResult(res).
		SetError(err).
		Get("/verify/{email}")
	if e != nil {
		return nil, &labstack.Error{
			Message: err.Error(),
		}
	}
	if labstack.IsError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}
