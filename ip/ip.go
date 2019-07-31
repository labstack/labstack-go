package ip

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go"
	"time"
)

type (
	Client struct {
		resty *resty.Client
	}

	Organization struct {
		Name string `json:"name"`
	}

	Flag struct {
		SVG   string `json:"svg"`
		PNG   string `json:"png"`
		Emoji string `json:"emoji"`
	}

	TimeZone struct {
		ID           string    `json:"id"`
		Name         string    `json:"name"`
		Abbreviation string    `json:"abbreviation"`
		Offset       int       `json:"offset"`
		Time         time.Time `json:"time"`
	}

	AS struct {
		Number       int    `json:"number"`
		Name         string `json:"name"`
		Organization string `json:"organization"`
	}

	LookupRequest struct {
		IP string
	}

	LookupResponse struct {
		IP           string        `json:"ip"`
		Hostname     string        `json:"hostname"`
		Version      string        `json:"version"`
		City         string        `json:"city"`
		Region       string        `json:"region"`
		RegionCode   string        `json:"region_code"`
		Postal       string        `json:"postal"`
		Country      string        `json:"country"`
		CountryCode  string        `json:"country_code"`
		Latitude     float64       `json:"latitude"`
		Longitude    float64       `json:"longitude"`
		Organization *Organization `json:"organization"`
		Flag         *Flag         `json:"flag"`
		Currencies   []string      `json:"currencies"`
		TimeZone     *TimeZone     `json:"time_zone"`
		Language     []string      `json:"languages"`
		AS           *AS           `json:"as"`
		Flags        []string      `json:"flags"`
	}
)

const (
	url = "https://ip.labstack.com/api/v1"
)

func New(key string) *Client {
	return &Client{
		resty: resty.New().SetHostURL(url).SetAuthToken(key),
	}
}

func (c *Client) Lookup(req *LookupRequest) (*LookupResponse, error) {
	res := new(LookupResponse)
	err := new(labstack.Error)
	r, e := c.resty.R().
		SetPathParams(map[string]string{
			"ip": req.IP,
		}).
		SetResult(res).
		SetError(err).
		Get("/{ip}")
	if e != nil {
		return nil, &labstack.Error{
			Message: e.Error(),
		}
	}
	if labstack.IsError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}
