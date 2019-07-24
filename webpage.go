package labstack

import (
	"github.com/labstack/labstack-go/webpage"
	"strconv"
)

func (c *Client) WebpageImage(req *webpage.ImageRequest) (*webpage.ImageResponse, *Error) {
	res := new(webpage.ImageResponse)
	err := new(Error)
	r, e := c.webpageResty.R().
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
		return nil, &Error{
			Message: err.Error(),
		}
	}
	if isError(r) {
		return nil, err
	}
	return res, nil
}

func (c *Client) WebpagePDF(req *webpage.PDFRequest) (*webpage.PDFResponse, *Error) {
	res := new(webpage.PDFResponse)
	err := new(Error)
	r, e := c.webpageResty.R().
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
		return nil, &Error{
			Message: err.Error(),
		}
	}
	if isError(r) {
		return nil, err
	}
	return res, nil
}
