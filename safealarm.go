package main

import (
	"log"
	"os"
)

// https://www.golang-book.com/books/intro

func main() {
	log.Print("STARTING")

	configFile := os.Args[1]
	config := NewConfiguration(configFile)

	message := NewMessage(config)
	message.SendStarting() // Send text message that the Pi is powering up

	doorLogic := NewDoorLogic(
		NewDoorSensor(),
		getIfEnabled(config.EnableDoorOpen, message.SendOpen),
		getIfEnabled(config.EnableDoorClosed, message.SendClosed),
		getIfEnabled(config.EnableDoorLeftOpen, message.SendForgot),
		message.Heartbeat,
		config.DoorOpenWaitSeconds,
		config.HeartbeatSeconds)
	doorLogic.Run()

	log.Print("EXITING")
}

func getIfEnabled(isEnabled bool, funcPtr func()) func() {
    if isEnabled {
        return funcPtr
    }
    return func() {}
}
