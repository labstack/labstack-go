package labstack

import "strconv"

type (
	ImageCompressRequest struct {
		File string
	}

	ImageCompressResponse struct {
		*Download
		Size int64 `json:"size"`
	}

	ImageResizeRequest struct {
		File   string
		Width  int
		Height int
		Format string
	}

	ImageResizeResponse struct {
		*Download
	}

	ImageWatermarkRequest struct {
		File     string
		Text     string
		Font     string
		Size     int
		Color    string
		Opacity  int
		Position string
		Margin   int
	}

	ImageWatermarkResponse struct {
		*Download
	}
)

func (c *Client) ImageCompress(req *ImageCompressRequest) (*ImageCompressResponse, *APIError) {
	res := new(ImageCompressResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/image/compress")
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

func (c *Client) ImageResize(req *ImageResizeRequest) (*ImageResizeResponse, *APIError) {
	res := new(ImageResizeResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetFile("file", req.File).
		SetFormData(map[string]string{
			"width":  strconv.Itoa(req.Width),
			"height": strconv.Itoa(req.Height),
			"format": req.Format,
		}).
		SetResult(res).
		SetError(err).
		Post("/image/resize")
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

func (c *Client) ImageWatermark(req *ImageWatermarkRequest) (*ImageWatermarkResponse, *APIError) {
	res := new(ImageWatermarkResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetFile("file", req.File).
		SetFormData(map[string]string{
			"text":     req.Text,
			"font":     req.Font,
			"size":     strconv.Itoa(req.Size),
			"color":    req.Color,
			"opacity":  strconv.Itoa(req.Opacity),
			"position": req.Position,
			"margin":   strconv.Itoa(req.Margin),
		}).
		SetResult(res).
		SetError(err).
		Post("/image/watermark")
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
