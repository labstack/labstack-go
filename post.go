package labstack

type (
	Post struct {
		*Client
	}

	PostVerifyResponse struct {
		ValidSyntax bool `json:"valid_syntax"`
		Deliverable bool `json:"deliverable"`
		InboxFull   bool `json:"inbox_full"`
		ValidDomain bool `json:"valid_domain"`
		Disposable  bool `json:"disposable"`
		CatchAll    bool `json:"catch_all"`
	}
)

func (p *Post) Verify(email string) (*PostVerifyResponse, *APIError) {
	res := new(PostVerifyResponse)
	err := new(APIError)
	r, e := p.resty.R().
		SetQueryParams(map[string]string{
			"email": email,
		}).
		SetResult(res).
		SetError(err).
		Get("/post/verify")
	if e != nil {
		return nil, &APIError{
			Message: e.Error(),
		}
	}
	if p.error(r) {
		return nil, err
	}
	return res, nil
}
