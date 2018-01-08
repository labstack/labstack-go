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

func (c *Client) PDFCompress(req *PDFCompressRequest) (*PDFCompressResponse, *APIError) {
	res := new(PDFCompressResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/pdf/compress")
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

func (c *Client) PDFImage(req *PDFImageRequest) (*PDFImageResponse, *APIError) {
	res := new(PDFImageResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetFile("file", req.File).
		SetFormData(map[string]string{
			"extract": strconv.FormatBool(req.Extract),
		}).
		SetResult(res).
		SetError(err).
		Post("/pdf/image")
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

func (c *Client) PDFSplit(req *PDFSplitRequest) (*PDFSplitResponse, *APIError) {
	res := new(PDFSplitResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetFile("file", req.File).
		SetFormData(map[string]string{
			"pages": req.Pages,
		}).
		SetResult(res).
		SetError(err).
		Post("/pdf/split")
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
