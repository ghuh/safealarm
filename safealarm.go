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
        func() {}, // Don't send a message on door close because it just wastes quota.  Also, a message will get printed in logs already.
        message.SendForgot,
        message.Heartbeat,
        config.DoorOpenWaitSeconds,
        config.HeartbeatSeconds)
    doorSensor.Run()

    log.Print("EXITING")
}
