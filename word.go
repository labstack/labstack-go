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

func (c *Client) WordLookup(req *WordLookupRequest) (res *WordLookupResponse, err *APIError) {
	res = new(WordLookupResponse)
	_, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/word/lookup")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
