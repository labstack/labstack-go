package labstack

import (
	"github.com/labstack/labstack-go/ip"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_IPLookup(t *testing.T) {
	res, err := client.IPLookup(&ip.LookupRequest{IP: "24.5.240.141"})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.Country)
	}
}
