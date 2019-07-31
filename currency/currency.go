package currency

import (
	"time"
)

type (
	Currency struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Symbol string `json:"symbol"`
	}

	ConvertRequest struct {
		Amount float64
		From   string
		To     string
	}

	ConvertResponse struct {
		Time   time.Time `json:"time"`
		Amount float64   `json:"amount"`
	}

	ListRequest struct {
	}

	ListResponse struct {
		Currencies []*Currency `json:"currencies"`
	}
)
