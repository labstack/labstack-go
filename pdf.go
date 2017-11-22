package labstack

type (
	PDFExtractImageRequest struct {
		File string
	}

	PDFExtractImageResponse struct {
		*Download
	}

	PDFToImageRequest struct {
		File string
	}

	PDFToImageResponse struct {
		*Download
	}
)

func (c *Client) PDFExtractImage(req *PDFExtractImageRequest) (res *PDFExtractImageResponse, err *APIError) {
	res = new(PDFExtractImageResponse)
	_, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/pdf/extract-image")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}

func (c *Client) PDFToImage(req *PDFToImageRequest) (res *PDFToImageResponse, err *APIError) {
	res = new(PDFToImageResponse)
	_, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/pdf/to-image")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
