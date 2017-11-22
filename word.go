package labstack

type (
	WordLookupRequest struct {
		Word string `json:"word"`
	}

	WordLookupResponse struct {
		Pronunciation []string
		Rhymes        []string
		Noun          []*WordLookupResult
		Verb          []*WordLookupResult
		Adverb        []*WordLookupResult
		Adjective     []*WordLookupResult
	}

	WordLookupResult struct {
		Definition string
		Synonyms   []string
		Antonyms   []string
		Examples   []string
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
