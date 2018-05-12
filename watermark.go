package labstack

import "strconv"

type (
	Watermark struct {
		*Client
	}

	WatermarkImageOptions struct {
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

func (w *Watermark) Image(file, text string, options WatermarkImageOptions) (*WatermarkImageResponse, *APIError) {
	res := new(WatermarkImageResponse)
	err := new(APIError)
	r, e := w.resty.R().
		SetFile("file", file).
		SetFormData(map[string]string{
			"text":     text,
			"font":     options.Font,
			"size":     strconv.Itoa(options.Size),
			"color":    options.Color,
			"opacity":  strconv.Itoa(options.Opacity),
			"position": options.Position,
			"margin":   strconv.Itoa(options.Margin),
		}).
		SetResult(res).
		SetError(err).
		Post("/watermark/image")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if w.error(r) {
		return nil, err
	}
	return res, nil
}
