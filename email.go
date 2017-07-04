package labstack

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/labstack/gommon/log"
)

type (
	// Email defines the LabStack email service.
	Email struct {
		sling  *sling.Sling
		logger *log.Logger
	}

	// EmailMessage defines the email message.
	EmailMessage struct {
		ID          string       `json:"id,omitempty"`
		From        string       `json:"from,omitempty"`
		To          string       `json:"to,omitempty"`
		Subject     string       `json:"subject,omitempty"`
		Body        string       `json:"body,omitempty"`
		Inlines     []*EmailFile `json:"inlines,omitempty"`
		Attachments []*EmailFile `json:"attachments,omitempty"`
	}

	// EmailFile defines the email message attachment/inline.
	EmailFile struct {
		// File name
		Name string `json:"name"`

		// File type
		Type string `json:"type"`

		// Base64 encoded file content
		Content string `json:"content"`
	}
)

// Email returns the email service.
func (c *Client) Email() *Email {
	return &Email{
		sling:  c.sling.Path("/email"),
		logger: c.logger,
	}
}

// Send sends the email message.
func (e *Email) Send(em *EmailMessage) (err error) {
	res, err := e.sling.Post("").BodyJSON(em).Receive(nil, nil)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("email: error sending email, status=%d, message=%v", res.StatusCode, err)
	}
	return
}
