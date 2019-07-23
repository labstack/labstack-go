package webpage

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/labstack-go"
)

type (
	Webpage struct {
		resty *resty.Client
	}

	WebpagePDFOptions struct {
		Layout string
		Format string
	}

	WebpagePDFResponse struct {
	}
)

func (w *Webpage) PDF(url string) (*WebpagePDFResponse, *labstack.APIError) {
	return w.PDFWithOptions(url, WebpagePDFOptions{})
}

func (w *Webpage) PDFWithOptions(url string, options WebpagePDFOptions) (*WebpagePDFResponse, *labstack.APIError) {
	res := new(WebpagePDFResponse)
	err := new(labstack.APIError)
	r, e := w.resty.R().
		SetQueryParams(map[string]string{
			"url":    url,
			"layout": options.Layout,
			"foramt": options.Format,
		}).
		SetResult(res).
		SetError(err).
		Get("/webpage/pdf")
	if e != nil {
		return nil, &labstack.APIError{
			Message: e.Error(),
		}
	}
	if w.Error(r) {
		return nil, err
	}
	return res, nil
}
