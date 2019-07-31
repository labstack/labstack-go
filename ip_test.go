package labstack

import (
	"github.com/labstack/labstack-go/ip"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Lookup(t *testing.T) {
	res, err := is.Lookup(&ip.LookupRequest{
		IP: "96.45.83.67",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.Country)
	}
}
