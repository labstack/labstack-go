package labstack

import (
	"github.com/labstack/labstack-go/webpage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_WebpageImage(t *testing.T) {
	res, err := client.WebpageImage(&webpage.ImageRequest{
		URL: "http://google.com",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.Image)
	}
}

func TestClient_WebpagePDF(t *testing.T) {
	res, err := client.WebpagePDF(&webpage.PDFRequest{
		URL: "http://google.com",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.PDF)
	}
}
