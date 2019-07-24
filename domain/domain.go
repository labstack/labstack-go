package domain

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go"
)

type (
	Client struct {
		resty *resty.Client
	}

	Record struct {
		Domain   string   `json:"domain"`
		Type     string   `json:"type"`
		Server   string   `json:"server"`
		A        string   `json:"a"`
		AAAA     string   `json:"aaaa"`
		CNAME    string   `json:"cname"`
		MX       string   `json:"mx"`
		NS       string   `json:"ns"`
		PTR      string   `json:"ptr"`
		Serial   uint32   `json:"serial"`
		Refresh  uint32   `json:"refresh"`
		Retry    uint32   `json:"retry"`
		Expire   uint32   `json:"expire"`
		Priority uint32   `json:"priority"`
		Weight   uint32   `json:"weight"`
		Port     uint32   `json:"port"`
		Target   string   `json:"target"`
		TXT      []string `json:"txt"`
		TTL      uint32   `json:"ttl"`
		Class    string   `json:"class"`
		SPF      []string `json:"spf"`
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
		CreatedDate string      `json:"created_date"`
		UpdatedDate string      `json:"updated_date"`
		ExpiryDate  string      `json:"expiry_date"`
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
