package main

import (
	"fmt"
	"log"

	"github.com/mailjet/mailjet-apiv3-go"
)

// https://icyapril.com/go/programming/2017/12/17/object-orientation-in-go.html
// For Verizon, send  to <number>@vtext.com to text a phone number

const numRetries int = 1

type Message struct {
	config Config
}

// NewMessage creates a new object for sending email messages.
func NewMessage(config Config) Message {
	message := Message{config}
	return message
}

// SendStarting sends a message that the program just started which also indicates that the machine just powered up.
func (m Message) SendStarting() {
	m.sendMessage("Safe Alarm Powering Up")
}

// SendOpen sends a message that the sensor is reading open.
func (m Message) SendOpen() {
	m.sendMessage("Safe Open")
}

// SendClosed sends a message that the sensor is reading closed.
func (m Message) SendClosed() {
	m.sendMessage("Safe Closed")
}

// SendForgot sends a message that the sensor was left open.
func (m Message) SendForgot() {
	m.sendMessage("Safe Left Open")
}

// Heartbeat sends a message that the program is still running
func (m Message) Heartbeat() {
	m.sendMessage("Safe Heartbeat")
}

// sendMessage sends the given message to all email addresses in the config file using mailjet.
// https://app.mailjet.com/transactional/sendapi
// https://github.com/mailjet/mailjet-apiv3-go
func (m Message) sendMessage(message string) {
	log.Printf("Sending message [ %v ]", message)

	// https://stackoverflow.com/a/38362784/10788820
	recipients := make([]mailjet.Recipient, len(m.config.TargetEmails))
	// https://stackoverflow.com/a/7782507/10788820
	for index := range m.config.TargetEmails {
		recipients[index] = mailjet.Recipient{
			Email: m.config.TargetEmails[index],
		}
	}

	email := &mailjet.InfoSendMail{
		FromEmail:  m.config.FromEmail,
		FromName:   m.config.FromName,
		Subject:    "",
		TextPart:   message,
		Recipients: recipients,
	}

	mailjetClient := mailjet.NewMailjetClient(m.config.Mailjet.PublicApiKey, m.config.Mailjet.PrivateApiKey)

	// Run in background thread so that retries don't hold up the execution of the program
	go func() {
		RunRetry(
			func() *string {
				res, err := mailjetClient.SendMail(email)

				if err != nil {
					msg := fmt.Sprintf("Error sending message: %v", err)
					return &msg
				}

				log.Printf("Successfully sent message: %v", res)
				return nil
			},
			-1)
	}()
}
