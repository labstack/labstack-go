package labstack

type (
	EmailVerifyRequest struct {
		Email string
	}

	EmailVerifyResponse struct {
		Syntax     bool   `json:"syntax"`
		Disposable bool   `json:"disposable"`
		Domain     bool   `json:"domain"`
		Mailbox    bool   `json:"mailbox"`
		Error      string `json:"error"`
	}
)

func (c *Client) EmailVerify(req *EmailVerifyRequest) (*EmailVerifyResponse, *APIError) {
	res := new(EmailVerifyResponse)
	err := new(APIError)
	r, e := c.resty.R().
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
	if c.error(r) {
		return nil, err
	}
	return res, nil
}
