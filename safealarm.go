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

    message.SendOpen()
}
