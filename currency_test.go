package labstack

import (
	"github.com/labstack/labstack-go/currency"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Convert(t *testing.T) {
	res, err := cs.Convert(&currency.ConvertRequest{
		Amount: 10,
		From:   "USD",
		To:     "INR",
	})
	if assert.Nil(t, err) {
		assert.NotZero(t, res.Amount)
	}
}

func TestClient_List(t *testing.T) {
	res, err := cs.List(&currency.ListRequest{})
	if assert.Nil(t, err) {
		assert.NotZero(t, len(res.Currencies))
	}
}
