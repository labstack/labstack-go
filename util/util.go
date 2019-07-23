package util

import "github.com/go-resty/resty/v2"

func Error(r *resty.Response) bool {
	return r.StatusCode() < 200 || r.StatusCode() >= 300
}
