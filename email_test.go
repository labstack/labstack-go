package labstack

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"io/ioutil"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	data, _ := ioutil.ReadFile("_fixture/logo.svg")
	ef := &EmailFile{
		Name:    "logo",
		Type:    "image/svg+xml",
		Content: base64.StdEncoding.EncodeToString(data),
	}
	em := &EmailMessage{
		From:    "Jack",
		To:      "jill@labstack.com",
		Subject: "Hello",
		Body:    "How are you doing?",
		Attachments: []*EmailFile{
			ef,
		},
		Inlines: []*EmailFile{
			ef,
		},
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		m := new(EmailMessage)
		if err := json.NewDecoder(r.Body).Decode(m); err == nil {
			if assert.EqualValues(t, em, m) {
				if assert.Len(t, m.Attachments, 1) && assert.Len(t, m.Inlines, 1) {
					if assert.EqualValues(t, m.Attachments[0], ef) && assert.EqualValues(t, m.Inlines[0], ef) {
						w.WriteHeader(http.StatusCreated)
						return
					}
				}
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()
	apiURL = ts.URL
	e := NewClient("").Email()
	assert.NoError(t, e.Send(em))
}
