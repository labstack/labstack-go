package labstack

import (
	"github.com/labstack/labstack-go/email"
	"github.com/labstack/labstack-go/util"
)

func (c *Client) EmailVerify(req *email.VerifyRequest) (*email.VerifyResponse, *Error) {
	res := new(email.VerifyResponse)
	ae := new(Error)
	r, err := c.emailResty.R().
		SetPathParams(map[string]string{
			"email": req.Email,
		}).
		SetResult(res).
		SetError(ae).
		Get("/verify/{email}")
	if err != nil {
		return nil, &Error{
			Message: err.Error(),
		}
	}
	if util.Error(r) {
		return nil, ae
	}
	return res, nil
}
