package labstack

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go/ip"
)

type (
	IPService struct {
		resty *resty.Client
	}
)

func (i *IPService) Lookup(req *ip.LookupRequest) (*ip.LookupResponse, error) {
	res := new(ip.LookupResponse)
	err := new(Error)
	r, e := i.resty.R().
		SetPathParams(map[string]string{
			"ip": req.IP,
		}).
		SetResult(res).
		SetError(err).
		Get("/{ip}")
	if e != nil {
		return nil, &Error{
			Message: e.Error(),
		}
	}
	if isError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}
