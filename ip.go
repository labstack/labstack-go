package labstack

import (
	"github.com/labstack/labstack-go/ip"
	"github.com/labstack/labstack-go/util"
)

func (c *Client) IPLookup(req *ip.LookupRequest) (*ip.LookupResponse, error) {
	res := new(ip.LookupResponse)
	err := new(Error)
	r, e := c.ipResty.R().
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
	if util.Error(r) {
		return nil, err
	}
	return res, nil
}
