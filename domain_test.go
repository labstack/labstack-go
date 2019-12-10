package labstack

import (
	"testing"

	"github.com/labstack/labstack-go/domain"
	"github.com/stretchr/testify/assert"
)

func TestClient_DNS(t *testing.T) {
	res, err := ds.DNS(&domain.DNSRequest{
		Type:   "A",
		Domain: "twilio.com",
	})
	if assert.Nil(t, err) {
		assert.NotZero(t, len(res.Records))
	}
}

func TestClient_Search(t *testing.T) {
	res, err := ds.Search(&domain.SearchRequest{
		Q: "twilio",
	})
	if assert.Nil(t, err) {
		assert.NotZero(t, len(res.Results))
	}
}

func TestClient_Status(t *testing.T) {
	res, err := ds.Status(&domain.StatusRequest{
		Domain: "twilio.com",
	})
	if assert.Nil(t, err) {
		assert.Equal(t, "unavailable", res.Result)
	}
}

func TestClient_Whois(t *testing.T) {
	res, err := ds.Whois(&domain.WhoisRequest{
		Domain: "twilio.com",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.Raw)
	}
}
