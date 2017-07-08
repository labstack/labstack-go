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
	file := &EmailFile{
		Name:    "logo",
		Type:    "image/svg+xml",
		Content: base64.StdEncoding.EncodeToString(data),
	}
	msg := &EmailMessage{
		From:    "Jack",
		To:      "jill@labstack.com",
		Subject: "Hello",
		Body:    "How are you doing?",
		Attachments: []*EmailFile{
			file,
		},
		Inlines: []*EmailFile{
			file,
		},
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		m := new(EmailMessage)
		if err := json.NewDecoder(r.Body).Decode(m); err == nil {
			if assert.EqualValues(t, msg, m) {
				if assert.Len(t, m.Attachments, 1) && assert.Len(t, m.Inlines, 1) {
					if assert.EqualValues(t, m.Attachments[0], file) && assert.EqualValues(t, m.Inlines[0], file) {
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
	assert.NoError(t, e.Send(msg))
}
