package labstack

import (
	"github.com/labstack/labstack-go/webpage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Image(t *testing.T) {
	res, err := ws.Image(&webpage.ImageRequest{
		URL: "amazon.com",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.Image)
	}
}

func TestClient_PDF(t *testing.T) {
	res, err := ws.PDF(&webpage.PDFRequest{
		URL: "amazon.com",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.PDF)
	}
}
