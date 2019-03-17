package main

import (
    "fmt"
    "github.com/mailjet/mailjet-apiv3-go"
)

// https://icyapril.com/go/programming/2017/12/17/object-orientation-in-go.html

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
    m.sendMessage("Sefe Left Open")
}

// https://app.mailjet.com/transactional/sendapi
func (m Message) sendMessage(message string) {
    mailjetClient := mailjet.NewMailjetClient(m.config.MJ.publicApiKey, m.config.MJ.privateApiKey)
    email := &mailjet.InfoSendMail {
      FromEmail: "kphayen@gmail.com",
      FromName: "Kevin Hayen",
      Subject: "",
      TextPart: message,
      Recipients: []mailjet.Recipient {
        mailjet.Recipient {
          Email: "7342553145@vtext.com",
        },
      },
    }
    res, err := mailjetClient.SendMail(email)
    if err != nil {
            fmt.Println(err)
    } else {
            fmt.Println("Success")
            fmt.Println(res)
    }
}
