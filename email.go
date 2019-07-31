package labstack

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go/email"
)

type (
	EmailService struct {
		resty *resty.Client
	}
)

func (es *EmailService) Verify(req *email.VerifyRequest) (*email.VerifyResponse, error) {
	res := new(email.VerifyResponse)
	err := new(Error)
	r, e := es.resty.R().
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
	if isError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}
