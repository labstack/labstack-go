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

func (c *Client) BarcodeGenerate(req *BarcodeGenerateRequest) (res *BarcodeGenerateResponse, err *APIError) {
	res = new(BarcodeGenerateResponse)
	_, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/barcode/generate")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
