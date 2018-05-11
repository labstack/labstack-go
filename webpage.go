package labstack

type (
	Webpage struct {
		*Client
	}

	WebpagePDFRequest struct {
		URL    string
		Layout string
		Format string
	}

	WebpagePDFResponse struct {
		*Download
	}
)

func (w *Webpage) PDF(req *WebpagePDFRequest) (*WebpagePDFResponse, *APIError) {
	res := new(WebpagePDFResponse)
	err := new(APIError)
	r, e := w.resty.R().
		SetQueryParams(map[string]string{
			"url":    req.URL,
			"layout": req.Layout,
			"foramt": req.Format,
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
