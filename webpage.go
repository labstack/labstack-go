package labstack

type (
	WebpagePDFRequest struct {
		URL    string `json:"url"`
		Size   string `json:"size"`
		Layout string `json:"width"`
	}

	WebpagePDFResponse struct {
		*Download
	}
)

func (c *Client) WebpagePDF(req *WebpagePDFRequest) (*WebpagePDFResponse, *APIError) {
	res := new(WebpagePDFResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/webpage/pdf")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if c.error(r) {
		return nil, err
	}
	return res, nil
}
