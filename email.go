package labstack

import (
	"fmt"

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
		From        string       `json:"from,omitempty"`
		To          string       `json:"to,omitempty"`
		Subject     string       `json:"subject,omitempty"`
		Body        string       `json:"body,omitempty"`
		Inlines     []*EmailFile `json:"inlines,omitempty"`
		Attachments []*EmailFile `json:"attachments,omitempty"`
	}

	// EmailFile defines the email message's attachment/inline.
	EmailFile struct {
		// File name
		Name string `json:"name"`

		// File type
		Type string `json:"type"`

		// Base64 encoded file content
		Content string `json:"content"`
	}

	// EmailStatus defines the email status.
	EmailStatus struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}

	// EmailError defines the email error.
	EmailError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

// Send sends the email message.
func (e *Email) Send(m *EmailMessage) (*EmailStatus, error) {
	es := new(EmailStatus)
	ee := new(EmailError)
	_, err := e.sling.Post("").BodyJSON(m).Receive(es, ee)
	if err != nil {
		return nil, err
	}
	return es, ee
}

func (e *EmailError) Error() string {
	return fmt.Sprintf("email error, code=%d, message=%s", e.Code, e.Message)
}
