package labstack

import (
	"github.com/labstack/labstack-go/ip"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_IPLookup(t *testing.T) {
	res, err := client.IPLookup(&ip.LookupRequest{
		IP: "96.45.83.67",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.Country)
	}
}
