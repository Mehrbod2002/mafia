package mailer

import (
	"fmt"
	"log"
)

// Message captures an outbound email.
type Message struct {
	To      string
	Subject string
	Body    string
}

// Sender defines a minimal email interface.
type Sender interface {
	Send(msg Message) error
}

// LoggerSender prints emails to stdout for development environments.
type LoggerSender struct{}

// Send outputs the message to the console.
func (LoggerSender) Send(msg Message) error {
	log.Printf("mail to=%s subject=%s body=%s", msg.To, msg.Subject, msg.Body)
	return nil
}

// MockSender collects sent messages for testing.
type MockSender struct {
	Sent []Message
}

// Send appends the message to Sent.
func (m *MockSender) Send(msg Message) error {
	m.Sent = append(m.Sent, msg)
	return nil
}

// Template formats a message body with simple placeholders.
func Template(template string, args ...interface{}) string {
	return fmt.Sprintf(template, args...)
}
