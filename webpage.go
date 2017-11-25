package labstack

type (
	WebpageToPDFRequest struct {
		URL    string `json:"url"`
		Size   string `json:"size"`
		Layout string `json:"width"`
	}

	WebpageToPDFResponse struct {
		*Download
	}
)

func (c *Client) WebpageToPDF(req *WebpageToPDFRequest) (res *WebpageToPDFResponse, err *APIError) {
	res = new(WebpageToPDFResponse)
	_, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/webpage/to-pdf")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
