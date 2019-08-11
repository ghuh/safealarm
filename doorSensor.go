package main

import (
	"log"

	"github.com/stianeikeland/go-rpio/v4"
)

// DoorSensor Object
type DoorSensor struct {
}

// pin is the GPIO pin on the board that we are reading from.
// Wire to 3.3 VDC Power and GPIO3 on P1 Pinout (26-pin Header)
// Raspberry Pi Model A pin guide: https://pi4j.com/1.2/pins/model-a-rev2.html
var (
	// Use mcu pin 22, corresponds to GPIO 3 on the pi
	pin = rpio.Pin(22)
)

// NewDoorSensor creates a new DoorSensor object to check if the door is open
func NewDoorSensor() DoorSensor {
	log.Print("Initializing GPIO")

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		log.Fatalf("Error opening GPIO pin: %v", err)
	}

	// Initialize pin to read
	pin.Input()

	doorSensor := DoorSensor{}
	return doorSensor
}

// IsOpen returns if the door is currently open
func (ds DoorSensor) IsOpen() bool {
	return pin.Read() == rpio.Low
}
