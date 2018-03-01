package labstack

import "strconv"

type (
	WatermarkImageRequest struct {
		File     string
		Text     string
		Font     string
		Size     int
		Color    string
		Opacity  int
		Position string
		Margin   int
	}

	WatermarkImageResponse struct {
		*Download
	}
)

func (c *Client) WatermarkImage(req *WatermarkImageRequest) (*WatermarkImageResponse, *APIError) {
	res := new(WatermarkImageResponse)
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
		Post("/watermark/image")
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
