package webpage

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go"
	"strconv"
)

type (
	Client struct {
		resty *resty.Client
	}

	ImageRequest struct {
		URL      string
		Language string
		TTL      int
		FullPage bool
		Retina   bool
		Width    int
		Height   int
		Delay    int
	}

	ImageResponse struct {
		Image       string `json:"image"`
		Cached      bool   `json:"cached"`
		Took        int    `json:"took"`
		GeneratedAt string `json:"generated_at"`
	}

	PDFRequest struct {
		URL         string
		Language    string
		TTL         int
		Size        string
		Width       int
		Height      int
		Orientation string
		Delay       int
	}

	PDFResponse struct {
		PDF         string `json:"pdf"`
		Cached      bool   `json:"cached"`
		Took        int    `json:"took"`
		GeneratedAt string `json:"generated_at"`
	}
)

const (
	url = "https://webpage.labstack.com/api/v1"
)

func New(key string) *Client {
	return &Client{
		resty: resty.New().SetHostURL(url).SetAuthToken(key),
	}
}

func (c *Client) Image(req *ImageRequest) (*ImageResponse, error) {
	res := new(ImageResponse)
	err := new(labstack.Error)
	r, e := c.resty.R().
		SetQueryParams(map[string]string{
			"url":       req.URL,
			"language":  req.Language,
			"ttl":       strconv.Itoa(req.TTL),
			"full_page": strconv.FormatBool(req.FullPage),
			"retina":    strconv.FormatBool(req.Retina),
			"width":     strconv.Itoa(req.Width),
			"height":    strconv.Itoa(req.Height),
			"delay":     strconv.Itoa(req.Delay),
		}).
		SetResult(res).
		SetError(err).
		Get("/image")
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

func (c *Client) PDF(req *PDFRequest) (*PDFResponse, error) {
	res := new(PDFResponse)
	err := new(labstack.Error)
	r, e := c.resty.R().
		SetQueryParams(map[string]string{
			"url":         req.URL,
			"language":    req.Language,
			"ttl":         strconv.Itoa(req.TTL),
			"size":        req.Size,
			"width":       strconv.Itoa(req.Width),
			"height":      strconv.Itoa(req.Height),
			"orientation": req.Orientation,
			"delay":       strconv.Itoa(req.Delay),
		}).
		SetResult(res).
		SetError(err).
		Get("/pdf")
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
