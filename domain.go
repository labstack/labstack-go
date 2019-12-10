package labstack

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go/domain"
)

type (
	DomainService struct {
		resty *resty.Client
	}
)

func (d *DomainService) DNS(req *domain.DNSRequest) (*domain.DNSResponse, error) {
	res := new(domain.DNSResponse)
	err := new(Error)
	r, e := d.resty.R().
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
	if isError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}

func (d *DomainService) Search(req *domain.SearchRequest) (*domain.SearchResponse, error) {
	res := new(domain.SearchResponse)
	err := new(Error)
	r, e := d.resty.R().
		SetQueryParam("q", req.Q).
		SetResult(res).
		SetError(err).
		Get("/search")
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

func (d *DomainService) Status(req *domain.StatusRequest) (*domain.StatusResponse, error) {
	res := new(domain.StatusResponse)
	err := new(Error)
	r, e := d.resty.R().
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
	if isError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}

func (d *DomainService) Whois(req *domain.WhoisRequest) (*domain.WhoisResponse, error) {
	res := new(domain.WhoisResponse)
	err := new(Error)
	r, e := d.resty.R().
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
	if isError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}
