package domain

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go"
	"time"
)

type (
	Client struct {
		resty *resty.Client
	}

	Record struct {
		Domain   string `json:"domain"`
		Type     string `json:"type"`
		Server   string `json:"server"`
		A        string
		AAAA     string
		CNAME    string
		MX       string
		NS       string
		PTR      string
		Serial   int    `json:"serial"`
		Refresh  int    `json:"refresh"`
		Retry    int    `json:"retry"`
		Expire   int    `json:"expire"`
		Priority int    `json:"priority"`
		Weight   int    `json:"weight"`
		Port     int    `json:"port"`
		Target   string `json:"target"`
		TXT      []string
		TTL      int    `json:"ttl"`
		Class    string `json:"class"`
		SPF      []string
	}

	Result struct {
		Domain string `json:"domain"`
		Zone   string `json:"zone"`
	}

	Registrar struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Url         string `json:"url"`
		WhoisServer string `json:"whois_server"`
	}

	Registrant struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Organization string `json:"organization"`
		Street       string `json:"street"`
		City         string `json:"city"`
		State        string `json:"state"`
		Zip          string `json:"zip"`
		Country      string `json:"country"`
		Phone        string `json:"phone"`
		Fax          string `json:"fax"`
		Email        string `json:"email"`
	}

	DNSRequest struct {
		Type   string
		Domain string
	}

	DNSResponse struct {
		Records []*Record
	}

	SearchRequest struct {
		Domain string
	}

	SearchResponse struct {
		Results []*Result
	}

	StatusRequest struct {
		Domain string
	}

	StatusResponse struct {
		Domain string   `json:"domain"`
		Zone   string   `json:"zone"`
		Result string   `json:"result"`
		Flags  []string `json:"flags"`
	}

	WhoisRequest struct {
		Domain string
	}

	WhoisResponse struct {
		Domain      string      `json:"domain"`
		Id          string      `json:"id"`
		Status      string      `json:"status"`
		CreatedDate time.Time   `json:"created_date"`
		UpdatedDate time.Time   `json:"updated_date"`
		ExpiryDate  time.Time   `json:"expiry_date"`
		NameServers []string    `json:"name_servers"`
		Dnssec      string      `json:"dnssec"`
		Registrar   *Registrar  `json:"registrar"`
		Registrant  *Registrant `json:"registrant"`
		Admin       *Registrant `json:"admin"`
		Technical   *Registrant `json:"technical"`
		Billing     *Registrant `json:"billing"`
		Raw         string      `json:"raw"`
	}
)

const (
	url = "https://domain.labstack.com/api/v1"
)

func New(key string) *Client {
	return &Client{
		resty: resty.New().SetHostURL(url).SetAuthToken(key),
	}
}

func (c *Client) DNS(req *DNSRequest) (*DNSResponse, error) {
	res := new(DNSResponse)
	err := new(labstack.Error)
	r, e := c.resty.R().
		SetPathParams(map[string]string{
			"type":   req.Type,
			"domain": req.Domain,
		}).
		SetResult(res).
		SetError(err).
		Get("/{type}/{domain}")
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

func (c *Client) Search(req *SearchRequest) (*SearchResponse, error) {
	res := new(SearchResponse)
	err := new(labstack.Error)
	r, e := c.resty.R().
		SetPathParams(map[string]string{
			"domain": req.Domain,
		}).
		SetResult(res).
		SetError(err).
		Get("/search/{domain}")
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

func (c *Client) Status(req *StatusRequest) (*StatusResponse, error) {
	res := new(StatusResponse)
	err := new(labstack.Error)
	r, e := c.resty.R().
		SetPathParams(map[string]string{
			"domain": req.Domain,
		}).
		SetResult(res).
		SetError(err).
		Get("/status/{domain}")
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

func (c *Client) Whois(req *WhoisRequest) (*WhoisResponse, error) {
	res := new(WhoisResponse)
	err := new(labstack.Error)
	r, e := c.resty.R().
		SetPathParams(map[string]string{
			"domain": req.Domain,
		}).
		SetResult(res).
		SetError(err).
		Get("/whois/{domain}")
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
