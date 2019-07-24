package labstack

import (
	"github.com/labstack/labstack-go/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_DomainDNS(t *testing.T) {
	res, err := client.DomainDNS(&domain.DNSRequest{Type: "A", Domain: "twillio.com"})
	if assert.Nil(t, err) {
		assert.NotZero(t, len(res.Records))
	}
}

func TestClient_DomainSearch(t *testing.T) {
	res, err := client.DomainSearch(&domain.SearchRequest{Domain: "twillio.com"})
	if assert.Nil(t, err) {
		assert.NotZero(t, len(res.Results))
	}
}

func TestClient_DomainStatus(t *testing.T) {
	res, err := client.DomainStatus(&domain.StatusRequest{Domain: "twillio.com"})
	if assert.Nil(t, err) {
		assert.Equal(t, "unavailable", res.Result)
	}
}

func TestClient_DomainWhois(t *testing.T) {
	res, err := client.DomainWhois(&domain.WhoisRequest{Domain: "twillio.com"})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, res.Raw)
	}
}
