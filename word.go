package labstack

type (
	WordLookupRequest struct {
		Word string `json:"word"`
	}

	WordLookupResponse struct {
		Pronunciation []string            `json:"pronunciation"`
		Rhymes        []string            `json:"rhymes"`
		Nouns         []*WordLookupResult `json:"nouns"`
		Verbs         []*WordLookupResult `json:"verbs"`
		Adverbs       []*WordLookupResult `json:"adverbs"`
		Adjectives    []*WordLookupResult `json:"adjectives"`
	}

	WordLookupResult struct {
		Definition string   `json:"definition"`
		Synonyms   []string `json:"synonyms"`
		Antonyms   []string `json:"antonyms"`
		Examples   []string `json:"examples"`
	}
)

func (c *Client) WordLookup(req *WordLookupRequest) (*WordLookupResponse, *APIError) {
	res := new(WordLookupResponse)
	err := new(APIError)
	r, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/word/lookup")
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
