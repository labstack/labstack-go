package labstack

import (
	"encoding/base64"
	"io/ioutil"
	"path/filepath"

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
		inlines     []string
		attachments []string
		Time        string       `json:"time,omitempty"`
		To          string       `json:"to,omitempty"`
		From        string       `json:"from,omitempty"`
		Subject     string       `json:"subject,omitempty"`
		Body        string       `json:"body,omitempty"`
		Inlines     []*emailFile `json:"inlines,omitempty"`
		Attachments []*emailFile `json:"attachments,omitempty"`
		Status      string       `json:"status,omitempty"`
	}

	emailFile struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Content string `json:"content"`
	}

	// EmailError defines the email error.
	EmailError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func NewEmailMessage(to, from, subject string) *EmailMessage {
	return &EmailMessage{
		To:      to,
		From:    from,
		Subject: subject,
	}
}

func (m *EmailMessage) addInlines() error {
	for _, inline := range m.inlines {
		data, err := ioutil.ReadFile(inline)
		if err != nil {
			return err
		}
		m.Inlines = append(m.Inlines, &emailFile{
			Name:    filepath.Base(inline),
			Content: base64.StdEncoding.EncodeToString(data),
		})
	}
	return nil
}

func (m *EmailMessage) addAttachments() error {
	for _, attachment := range m.attachments {
		data, err := ioutil.ReadFile(attachment)
		if err != nil {
			return err
		}
		m.Inlines = append(m.Attachments, &emailFile{
			Name:    filepath.Base(attachment),
			Content: base64.StdEncoding.EncodeToString(data),
		})
	}
	return nil
}

func (m *EmailMessage) AddInline(path string) {
	m.inlines = append(m.inlines, path)
}

func (m *EmailMessage) AddAttachment(path string) {
	m.attachments = append(m.attachments, path)
}

// Send sends the email message.
func (e *Email) Send(m *EmailMessage) (*EmailMessage, error) {
	if err := m.addInlines(); err != nil {
		return nil, err
	}
	if err := m.addAttachments(); err != nil {
		return nil, err
	}
	em := new(EmailMessage)
	ee := new(EmailError)
	_, err := e.sling.Post("").BodyJSON(m).Receive(em, ee)
	if err != nil {
		return nil, err
	}
	if ee.Code == 0 {
		return em, nil
	}
	return nil, ee
}

func (e *EmailError) Error() string {
	return e.Message
}
