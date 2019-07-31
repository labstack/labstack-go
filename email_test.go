package labstack

import (
	"github.com/labstack/labstack-go/email"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Verify(t *testing.T) {
	res, err := es.Verify(&email.VerifyRequest{
		Email: "jon@labstack.com",
	})
	if assert.Nil(t, err) {
		assert.Equal(t, "deliverable", res.Result)
	}
}
