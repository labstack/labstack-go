package labstack

type (
	Webpage struct {
		*Client
	}

	WebpagePDFOptions struct {
		Layout string
		Format string
	}

	WebpagePDFResponse struct {
		*Download
	}
)

func (w *Webpage) PDF(url string, options WebpagePDFOptions) (*WebpagePDFResponse, *APIError) {
	res := new(WebpagePDFResponse)
	err := new(APIError)
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
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if w.error(r) {
		return nil, err
	}
	return res, nil
}
