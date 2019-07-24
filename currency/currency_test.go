package currency

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	client = New(os.Getenv("KEY"))
)

func TestClient_CurrencyConvert(t *testing.T) {
	res, err := client.Convert(&ConvertRequest{
		Amount: 10,
		From:   "USD",
		To:     "INR",
	})
	if assert.Nil(t, err) {
		assert.NotZero(t, res.Amount)
	}
}

func TestClient_CurrencyList(t *testing.T) {
	res, err := client.List(&ListRequest{})
	if assert.Nil(t, err) {
		assert.NotZero(t, len(res.Currencies))
	}
}
