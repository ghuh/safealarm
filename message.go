package main

import (
    "log"
)

// https://icyapril.com/go/programming/2017/12/17/object-orientation-in-go.html
// For Verizon, send  to <number>@vtext.com to text a phone number

type Message struct {
    config Config
}

func NewMessage(config Config) Message {
    message := Message{config}
    return message
}

func (m Message) SendOpen() {
    m.sendMessage("Safe Open")
}

func (m Message) SendClosed() {
    m.sendMessage("Safe Closed")
}

func (m Message) SendForgot() {
    m.sendMessage("Safe Left Open")
}

func (m Message) Heartbeat() {
    m.sendMessage("Safe Heartbeat")
}

// https://app.mailjet.com/transactional/sendapi
func (m Message) sendMessage(message string) {
    log.Printf("Sending message [ %v ]", message)

    /* TODO
    // https://stackoverflow.com/a/38362784/10788820
    recipients := make([]mailjet.Recipient, len(m.config.targetEmails))
    // https://stackoverflow.com/a/7782507/10788820
    for index, _ := range m.config.targetEmails {
        recipients[index] = mailjet.Recipient{
            Email: m.config.targetEmails[index],
        }
    }

    email := &mailjet.InfoSendMail {
      FromEmail: "kphayen@gmail.com",
      FromName: "Kevin Hayen",
      Subject: "",
      TextPart: message,
      Recipients: recipients,
    }

    mailjetClient := mailjet.NewMailjetClient(m.config.Mailjet.publicApiKey, m.config.Mailjet.privateApiKey)

    res, err := mailjetClient.SendMail(email)
    if err != nil {
        log.Printf("Error sending message: %v", err)
    } else {
        log.Printf("Successfully sent message: %v", res)
    }
    */
}
