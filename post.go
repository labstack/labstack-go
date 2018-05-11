package labstack

type (
	Post struct {
		*Client
	}

	PostVerifyRequest struct {
		Email string
	}

	PostVerifyResponse struct {
		Syntax     bool   `json:"syntax"`
		Disposable bool   `json:"disposable"`
		Domain     bool   `json:"domain"`
		Mailbox    bool   `json:"mailbox"`
		Error      string `json:"error"`
	}
)

func (p *Post) Verify(req *PostVerifyRequest) (*PostVerifyResponse, *APIError) {
	res := new(PostVerifyResponse)
	err := new(APIError)
	r, e := p.resty.R().
		SetQueryParams(map[string]string{
			"email": req.Email,
		}).
		SetResult(res).
		SetError(err).
		Get("/email/verify")
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
