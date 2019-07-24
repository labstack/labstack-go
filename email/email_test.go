package email

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	client = New(os.Getenv("KEY"))
)

func TestClient_EmailVerify(t *testing.T) {
	res, err := client.Verify(&VerifyRequest{
		Email: "jon@labstack.com",
	})
	if assert.Nil(t, err) {
		assert.Equal(t, "deliverable", res.Result)
	}
}
