package labstack

import (
	"github.com/labstack/labstack-go/currency"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_CurrencyConvert(t *testing.T) {
	res, err := client.CurrencyConvert(&currency.ConvertRequest{
		Amount: 10,
		From:   "USD",
		To:     "INR",
	})
	if assert.Nil(t, err) {
		assert.NotZero(t, res.Amount)
	}
}

func TestClient_CurrencyList(t *testing.T) {
	res, err := client.CurrencyList(&currency.ListRequest{})
	if assert.Nil(t, err) {
		assert.NotZero(t, len(res.Currencies))
	}
}
