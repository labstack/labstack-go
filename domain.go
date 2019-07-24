package labstack

import (
	"github.com/labstack/labstack-go/domain"
)

func (c *Client) DomainDNS(req *domain.DNSRequest) (*domain.DNSResponse, *Error) {
	res := new(domain.DNSResponse)
	err := new(Error)
	r, e := c.domainResty.R().
		SetPathParams(map[string]string{
			"type":   req.Type,
			"domain": req.Domain,
		}).
		SetResult(res).
		SetError(err).
		Get("/{type}/{domain}")
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

func (c *Client) DomainSearch(req *domain.SearchRequest) (*domain.SearchResponse, *Error) {
	res := new(domain.SearchResponse)
	err := new(Error)
	r, e := c.domainResty.R().
		SetPathParams(map[string]string{
			"domain": req.Domain,
		}).
		SetResult(res).
		SetError(err).
		Get("/search/{domain}")
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

func (c *Client) DomainStatus(req *domain.StatusRequest) (*domain.StatusResponse, *Error) {
	res := new(domain.StatusResponse)
	err := new(Error)
	r, e := c.domainResty.R().
		SetPathParams(map[string]string{
			"domain": req.Domain,
		}).
		SetResult(res).
		SetError(err).
		Get("/status/{domain}")
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

func (c *Client) DomainWhois(req *domain.WhoisRequest) (*domain.WhoisResponse, *Error) {
	res := new(domain.WhoisResponse)
	err := new(Error)
	r, e := c.domainResty.R().
		SetPathParams(map[string]string{
			"domain": req.Domain,
		}).
		SetResult(res).
		SetError(err).
		Get("/whois/{domain}")
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
