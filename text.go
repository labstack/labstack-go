package labstack

type (
	TextSummaryRequest struct {
		Text     string `json:"text"`
		URL      string `json:"url"`
		Language string `json:"language"`
		Length   int    `json:"length"`
	}

	TextSummaryResponse struct {
		Summary string `json:"summary"`
	}

	TextSentimentRequest struct {
		Text string `json:"text"`
	}

	TextSentimentResponse struct {
		Subjectivity float32 `json:"subjectivity"`
		Polarity     float32 `json:"polarity"`
	}

	TextSpellCheckRequest struct {
		Text string `json:"text"`
	}

	TextSpellCheckResponse struct {
		Misspelled []*TextSpellCheckMisspelled `json:"misspelled"`
	}

	TextSpellCheckMisspelled struct {
		Word        string   `json:"word"`
		Offset      int      `json:"offset"`
		Suggestions []string `json:"suggestions"`
	}
)

func (c *Client) TextSummary(req *TextSummaryRequest) (res *TextSummaryResponse, err *APIError) {
	res = new(TextSummaryResponse)
	_, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/text/summary")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}

func (c *Client) TextSentiment(req *TextSentimentRequest) (res *TextSentimentResponse, err *APIError) {
	res = new(TextSentimentResponse)
	_, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/text/sentiment")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}

func (c *Client) TextSpellCheck(req *TextSpellCheckRequest) (res *TextSpellCheckResponse, err *APIError) {
	res = new(TextSpellCheckResponse)
	_, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/text/spell-check")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
