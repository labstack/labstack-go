package domain

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	client = New(os.Getenv("KEY"))
)

func TestClient_DNS(t *testing.T) {
	res, err := client.DNS(&DNSRequest{
		Type:   "A",
		Domain: "twilio.com",
	})
	if assert.Nil(t, err) {
		assert.NotZero(t, len(res.Records))
	}
}

func TestClient_Search(t *testing.T) {
	res, err := client.Search(&SearchRequest{
		Domain: "twilio.com",
	})
	if assert.Nil(t, err) {
		assert.NotZero(t, len(res.Results))
	}
}

func TestClient_Status(t *testing.T) {
	res, err := client.Status(&StatusRequest{
		Domain: "twilio.com",
	})
	if assert.Nil(t, err) {
		assert.Equal(t, "unavailable", res.Result)
	}
}

func TestClient_Whois(t *testing.T) {
	res, err := client.Whois(&WhoisRequest{
		Domain: "twilio.com",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.Raw)
	}
}
