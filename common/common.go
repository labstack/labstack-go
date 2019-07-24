package common

import "github.com/go-resty/resty/v2"

type (
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func IsError(r *resty.Response) bool {
	return r.StatusCode() < 200 || r.StatusCode() >= 300
}

func (e *Error) Error() string {
	return e.Message
}
