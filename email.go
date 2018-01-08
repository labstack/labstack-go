package labstack

type (
	EmailVerifyRequest struct {
		Email string `json:"email"`
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
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/email/verify")
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
