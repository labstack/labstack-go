package ip

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	client = New(os.Getenv("KEY"))
)

func TestClient_Lookup(t *testing.T) {
	res, err := client.Lookup(&LookupRequest{
		IP: "96.45.83.67",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.Country)
	}
}
