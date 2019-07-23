package domain

import "github.com/go-resty/resty/v2"

type (
	Domain struct {
		resty *resty.Client
	}
)
