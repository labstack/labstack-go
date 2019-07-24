package labstack

import (
	"github.com/labstack/labstack-go/email"
)

func (c *Client) EmailVerify(req *email.VerifyRequest) (*email.VerifyResponse, *Error) {
	res := new(email.VerifyResponse)
	err := new(Error)
	r, e := c.emailResty.R().
		SetPathParams(map[string]string{
			"email": req.Email,
		}).
		SetResult(res).
		SetError(err).
		Get("/verify/{email}")
	if e != nil {
		return nil, &Error{
			Message: err.Error(),
		}
	}
	if isError(r) {
		return nil, err
	}
	return res, nil
}
