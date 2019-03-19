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

    doorSensor := NewDoorSensor( message.SendOpen, message.SendClosed, message.SendForgot, config.DoorOpenWaitSeconds)
    doorSensor.Run()

    log.Print("EXITING")
}
