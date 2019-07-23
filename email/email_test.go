package email

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	key = os.Getenv("KEY")
)

func TestEmail_Verify(t *testing.T) {
	e := New()
	res, err := e.Verify(&VerifyRequest{Email: "vr@labstack.com"})
	if assert.Nil(t, err) {
		assert.Equal(t, "deliverable", res.Result)
	}
}
