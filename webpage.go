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

func (c *Client) WebpagePDF(req *WebpagePDFRequest) (res *WebpagePDFResponse, err *APIError) {
	res = new(WebpagePDFResponse)
	_, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/webpage/pdf")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
