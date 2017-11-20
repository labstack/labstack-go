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

func (c *Client) EmailVerify(req *EmailVerifyRequest) (res *EmailVerifyResponse, err *APIError) {
	res = new(EmailVerifyResponse)
	_, e := c.resty.R().
		SetBody(req).
		SetResult(res).
		SetError(err).
		Post("/email/verify")
	if e != nil {
		err = new(APIError)
		err.Message = e.Error()
	}
	return
}
