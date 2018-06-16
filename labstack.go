package labstack

type (
	Map map[string]interface{}

	APIError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)
