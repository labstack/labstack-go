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

	TextSpellcheckRequest struct {
		Text string `json:"text"`
	}

	TextSpellcheckResponse struct {
		Misspelled []*TextSpellcheckMisspelled `json:"misspelled"`
	}

	TextSpellcheckMisspelled struct {
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

func (c *Client) TextSpellcheck(req *TextSpellcheckRequest) (res *TextSpellcheckResponse, err *APIError) {
	res = new(TextSpellcheckResponse)
	_, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/text/spellcheck")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
