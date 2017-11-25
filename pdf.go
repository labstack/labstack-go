package labstack

type (
	PDFImageRequest struct {
		File string
	}

	PDFImageResponse struct {
		*Download
	}
)

func (c *Client) PDFImage(req *PDFImageRequest) (res *PDFImageResponse, err *APIError) {
	res = new(PDFImageResponse)
	_, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/pdf/image")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
