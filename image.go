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

func (c *Client) ImageCompress(req *ImageCompressRequest) (res *ImageCompressResponse, err *APIError) {
	res = new(ImageCompressResponse)
	_, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/image/compress")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}

func (c *Client) ImageResize(req *ImageResizeRequest) (res *ImageResizeResponse, err *APIError) {
	res = new(ImageResizeResponse)
	_, e := c.resty.R().
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
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}

func (c *Client) ImageWatermark(req *ImageWatermarkRequest) (res *ImageWatermarkResponse, err *APIError) {
	res = new(ImageWatermarkResponse)
	_, e := c.resty.R().
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
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
