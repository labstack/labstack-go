package email

type (
	VerifyRequest struct {
		Email string
	}

	VerifyResponse struct {
		Email    string   `json:"email"`
		Username string   `json:"username"`
		Domain   string   `json:"domain"`
		Result   string   `json:"result"`
		Flags    []string `json:"flags"`
	}
)
