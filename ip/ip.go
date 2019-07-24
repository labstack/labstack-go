package ip

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go"
)

type (
	Client struct {
		resty *resty.Client
	}

	Currency struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Symbol string `json:"symbol"`
	}

	Organization struct {
		Name string `json:"name"`
	}

	Flag struct {
		Image        string `json:"image"`
		Emoji        string `json:"emoji"`
		EmojiUnicode string `json:"emoji_unicode"`
	}

	TimeZone struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation"`
		Offset       int32  `json:"offset"`
		Time         string `json:"time"`
	}

	Language struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}

	AS struct {
		Number       int64  `json:"number"`
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
		Currency     *Currency     `json:"currency"`
		TimeZone     *TimeZone     `json:"time_zone"`
		Language     []*Language   `json:"languages"`
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
