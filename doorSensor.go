package main

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
)

// Raspberry Pi Model A pin guide: https://pi4j.com/1.2/pins/model-a-rev2.html
// Wire to 3.3 VDC Power and GPIO3 on P1 Pinout (26-pin Header)

// https://github.com/stianeikeland/go-rpio/blob/master/examples/event/event.go

type DoorSensor struct {
	onOpen func()
	onClose func()
	onForgot func()
	doorOpenWaitDuration time.Duration

	isOpen bool
}

var (
	// Use mcu pin 22, corresponds to GPIO 3 on the pi
	pin = rpio.Pin(22)
)

func NewDoorSensor( onOpen func(), onClose func(), onForgot func(), doorOpenWaitSeconds int32 ) DoorSensor {
	doorOpenWaitDuration, _ := time.ParseDuration(string(doorOpenWaitSeconds) + "s")
	doorSensor := DoorSensor{onOpen, onClose, onForgot, doorOpenWaitDuration, false}
	return doorSensor
}

func (ds DoorSensor) Run() {
	log.Print("Initializing GPIO")

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		log.Fatalf("Error opening GPIO pin: %v", err)
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	pin.Input()

	ds.isOpen = pin.Read() == rpio.Low
	ds.printState()

	log.Print("Listening for door sensor")

	currentForgotTime := new(time.Time)

	for true {
		// Get new state of sensor
		currentlyOpen := pin.Read() == rpio.Low

		// Send forgot alarm if open for too long
		if currentForgotTime != nil && time.Now().After(*currentForgotTime) {
			currentForgotTime = nil
			ds.onForgot()
		}

		// Check if state has changed
		if currentlyOpen != ds.isOpen {
			ds.isOpen = currentlyOpen
			ds.printState()
			if (ds.isOpen) {
				// Need to get reference
				newForgotTime := time.Now().Add(ds.doorOpenWaitDuration)
				currentForgotTime = &newForgotTime

				ds.onOpen()
			} else {
				currentForgotTime = nil
				ds.onClose()
			}
		}

		// Sleep half second before next loop
		time.Sleep(time.Second / 2)
	}

	log.Print("Done listening for door sensor")
}

func (ds DoorSensor) printState() {
	if ds.isOpen {
		log.Print("Door open")
	} else {
		log.Print("Door closed")
	}
}
