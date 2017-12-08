package labstack

import "strconv"

type (
	PDFImageRequest struct {
		File    string
		Extract bool
	}

	PDFImageResponse struct {
		*Download
	}

	PDFSplitRequest struct {
		File  string
		Pages string
	}

	PDFSplitResponse struct {
		*Download
	}

	PDFCompressRequest struct {
		File string
	}

	PDFCompressResponse struct {
		*Download
		Size int64 `json:"size"`
	}
)

func (c *Client) PDFCompress(req *PDFCompressRequest) (res *PDFCompressResponse, err *APIError) {
	res = new(PDFCompressResponse)
	_, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/pdf/compress")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}

func (c *Client) PDFImage(req *PDFImageRequest) (res *PDFImageResponse, err *APIError) {
	res = new(PDFImageResponse)
	_, e := c.resty.R().
		SetFile("file", req.File).
		SetFormData(map[string]string{
			"extract": strconv.FormatBool(req.Extract),
		}).
		SetResult(res).
		SetError(err).
		Post("/pdf/image")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}

func (c *Client) PDFSplit(req *PDFSplitRequest) (res *PDFSplitResponse, err *APIError) {
	res = new(PDFSplitResponse)
	_, e := c.resty.R().
		SetFile("file", req.File).
		SetFormData(map[string]string{
			"pages": req.Pages,
		}).
		SetResult(res).
		SetError(err).
		Post("/pdf/split")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
