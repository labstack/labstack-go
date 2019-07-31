package labstack

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go/webpage"
	"strconv"
)

type (
	WebpageService struct {
		resty *resty.Client
	}
)

func (w *WebpageService) Image(req *webpage.ImageRequest) (*webpage.ImageResponse, error) {
	res := new(webpage.ImageResponse)
	err := new(Error)
	r, e := w.resty.R().
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
	if isError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}

func (w *WebpageService) PDF(req *webpage.PDFRequest) (*webpage.PDFResponse, error) {
	res := new(webpage.PDFResponse)
	err := new(Error)
	r, e := w.resty.R().
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
	if isError(r.StatusCode()) {
		return nil, err
	}
	return res, nil
}
