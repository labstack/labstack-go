package labstack

type (
	Email struct {
		*Client
	}

	EmailVerifyResponse struct {
		ValidSyntax bool `json:"valid_syntax"`
		Deliverable bool `json:"deliverable"`
		InboxFull   bool `json:"inbox_full"`
		ValidDomain bool `json:"valid_domain"`
		Disposable  bool `json:"disposable"`
		CatchAll    bool `json:"catch_all"`
	}
)

func (e *Email) Verify(email string) (*EmailVerifyResponse, *APIError) {
	res := new(EmailVerifyResponse)
	ae := new(APIError)
	r, err := e.resty.R().
		SetQueryParams(map[string]string{
			"email": email,
		}).
		SetResult(res).
		SetError(ae).
		Get("/post/verify")
	if err != nil {
		return nil, &APIError{
			Message: err.Error(),
		}
	}
	if e.error(r) {
		return nil, ae
	}
	return res, nil
}
