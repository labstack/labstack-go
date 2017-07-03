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

	// Message defines the email message.
	Message struct {
		ID          string  `json:"id,omitempty"`
		From        string  `json:"from,omitempty"`
		To          string  `json:"to,omitempty"`
		Subject     string  `json:"subject,omitempty"`
		Body        string  `json:"body,omitempty"`
		Inlines     []*File `json:"inlines,omitempty"`
		Attachments []*File `json:"attachments,omitempty"`
	}

	// File defines the email message attachment/inline.
	File struct {
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
func (e *Email) Send(m *Message) (err error) {
	res, err := e.sling.Post("").BodyJSON(m).Receive(nil, nil)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("email: error sending email, status=%d, message=%v", res.StatusCode, err)
	}
	return
}
