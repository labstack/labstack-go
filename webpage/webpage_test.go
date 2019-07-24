package webpage

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	client = New(os.Getenv("KEY"))
)

func TestClient_Image(t *testing.T) {
	res, err := client.Image(&ImageRequest{
		URL: "amazon.com",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.Image)
	}
}

func TestClient_PDF(t *testing.T) {
	res, err := client.PDF(&PDFRequest{
		URL: "amazon.com",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.PDF)
	}
}
