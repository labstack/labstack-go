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
		Width  int  `form:"width"`
		Height int  `form:"height"`
		Crop   bool `form:"crop"`
	}

	ImageResizeResponse struct {
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
			"crop":   strconv.FormatBool(req.Crop),
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
