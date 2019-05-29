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

    doorSensor := NewDoorSensor(
        message.SendOpen,
        message.SendClosed,
        message.SendForgot,
        message.Heartbeat,
        config.DoorOpenWaitSeconds,
        config.HeartbeatSeconds)
    doorSensor.Run()

    log.Print("EXITING")
}
