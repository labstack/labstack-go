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

func (c *Client) TextSummary(req *TextSummaryRequest) (*TextSummaryResponse, *APIError) {
	res := new(TextSummaryResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/text/summary")
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

func (c *Client) TextSentiment(req *TextSentimentRequest) (*TextSentimentResponse, *APIError) {
	res := new(TextSentimentResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/text/sentiment")
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

func (c *Client) TextSpellcheck(req *TextSpellcheckRequest) (*TextSpellcheckResponse, *APIError) {
	res := new(TextSpellcheckResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/text/spellcheck")
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
