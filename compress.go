package labstack

type (
	CompressAudioRequest struct {
		File string
	}

	CompressAudioResponse struct {
		*Download
		Size int64 `json:"size"`
	}

	CompressImageRequest struct {
		File string
	}

	CompressImageResponse struct {
		*Download
		Size int64 `json:"size"`
	}

	CompressPDFRequest struct {
		File string
	}

	CompressPDFResponse struct {
		*Download
		Size int64 `json:"size"`
	}

	CompressVideoRequest struct {
		File string
	}

	CompressVideoResponse struct {
		*Download
		Size int64 `json:"size"`
	}
)

func (c *Client) CompressAudio(req *CompressAudioRequest) (*CompressAudioResponse, *APIError) {
	res := new(CompressAudioResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/compress/audio")
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

func (c *Client) CompressImage(req *CompressImageRequest) (*CompressImageResponse, *APIError) {
	res := new(CompressImageResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/compress/image")
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

func (c *Client) CompressPDF(req *CompressPDFRequest) (*CompressPDFResponse, *APIError) {
	res := new(CompressPDFResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/compress/pdf")
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

func (c *Client) CompressVideo(req *CompressVideoRequest) (*CompressVideoResponse, *APIError) {
	res := new(CompressVideoResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/compress/video")
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
