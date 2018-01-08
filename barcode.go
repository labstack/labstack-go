package labstack

type (
	BarcodeGenerateRequest struct {
		Format  string `json:"format"`
		Content string `json:"content"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
	}

	BarcodeGenerateResponse struct {
		*Download
	}

	BarcodeScanRequest struct {
		File string
	}

	BarcodeScanResponse struct {
		Format      string `json:"format"`
		Content     string `json:"content"`
		ContentType string `json:"content_type"`
	}
)

func (c *Client) BarcodeGenerate(req *BarcodeGenerateRequest) (*BarcodeGenerateResponse, *APIError) {
	res := new(BarcodeGenerateResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/barcode/generate")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if success(r) {
		return res, nil
	}
	return nil, err
}

func (c *Client) BarcodeScan(req *BarcodeScanRequest) (*BarcodeScanResponse, *APIError) {
	res := new(BarcodeScanResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetFile("file", req.File).
		SetResult(res).
		SetError(err).
		Post("/barcode/scan")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if success(r) {
		return res, nil
	}
	return nil, err
}
