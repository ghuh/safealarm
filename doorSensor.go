package main

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
)

// Raspberry Pi Model A pin guide: https://pi4j.com/1.2/pins/model-a-rev2.html
// Pin 22 mapping in code: https://github.com/stianeikeland/go-rpio/blob/master/rpio.go#L49
// https://github.com/stianeikeland/go-rpio/blob/master/examples/event/event.go

type DoorSensor struct {
	onOpen func()
	onClose func()
	isOpen bool
}

var (
	// Use mcu pin 22, corresponds to GPIO 3 on the pi
	pin = rpio.Pin(22)
)

func NewDoorSensor(onOpen func(), onClose func()) DoorSensor {
	doorSensor := DoorSensor{onOpen, onClose, false}
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

	pin.Detect(rpio.FallEdge) // enable falling edge event detection
	pin.Detect(rpio.RiseEdge) // enable falling edge event detection

	log.Print("Listening for door sensor")

	for true {
		if pin.EdgeDetected() { // check if event occurred
			ds.isOpen = !ds.isOpen
			ds.printState()
			if (ds.isOpen) {
				ds.onOpen()
			} else {
				ds.onClose()
			}
		}
		time.Sleep(time.Second / 2)
	}

	log.Print("Done listening for door sensor")

	pin.Detect(rpio.NoEdge) // disable edge event detection
}

func (ds DoorSensor) printState() {
	if ds.isOpen {
		log.Print("Door open")
	} else {
		log.Print("Door closed")
	}
}
