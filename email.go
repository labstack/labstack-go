package labstack

import (
	"fmt"

	"encoding/base64"
	"io/ioutil"
	"mime"
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
		Time        string       `json:"time,omitempty"`
		From        string       `json:"from,omitempty"`
		To          string       `json:"to,omitempty"`
		Subject     string       `json:"subject,omitempty"`
		Body        string       `json:"body,omitempty"`
		Status      string       `json:"status,omitempty"`
		Inlines     []*emailFile `json:"inlines,omitempty"`
		Attachments []*emailFile `json:"attachments,omitempty"`
		inlines     []string
		attachments []string
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

func emailFileFromPath(path string) (*emailFile, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &emailFile{
		Name:    filepath.Base(path),
		Type:    mime.TypeByExtension(filepath.Ext(path)),
		Content: base64.StdEncoding.EncodeToString(data),
	}, nil
}

func (m *EmailMessage) addFiles() error {
	for _, path := range m.inlines {
		file, err := emailFileFromPath(path)
		if err != nil {
			return err
		}
		m.Inlines = append(m.Inlines, file)
	}
	for _, path := range m.attachments {
		file, err := emailFileFromPath(path)
		if err != nil {
			return err
		}
		m.Attachments = append(m.Attachments, file)
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
	if err := m.addFiles(); err != nil {
		return nil, err
	}
	em := new(EmailMessage)
	ee := new(EmailError)
	_, err := e.sling.Post("").BodyJSON(m).Receive(em, ee)
	if err != nil {
		return nil, err
	}
	return em, ee
}

func (e *EmailError) Error() string {
	return fmt.Sprintf("email error, code=%d, message=%s", e.Code, e.Message)
}
