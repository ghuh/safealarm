package main

import (
	"flag"
	"log"
)

// https://www.golang-book.com/books/intro

func main() {
    log.Print("STARTING")

    var configFile string
	flag.StringVar(&configFile,"CONFIG_FILE", "", "Config file path")

    config := NewConfiguration(configFile)

    message := NewMessage(config)

    message.SendOpen()
}
