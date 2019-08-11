package main

import (
	"log"
	"strconv"
	"time"
)

// DoorLogic object
type DoorLogic struct {
	doorSensor           DoorSensor
	onOpen               func()
	onClose              func()
	onForgot             func()
	onHeartbeat          func()
	doorOpenWaitDuration time.Duration
	heartbeatSeconds     int

	isOpen bool
}

// NewDoorLogic creates a new DoorLogic object that you Run() and then it'll fire callbacks on events to the door.
func NewDoorLogic(
	doorSensor DoorSensor,
	onOpen func(),
	onClose func(),
	onForgot func(),
	onHeartbeat func(),
	doorOpenWaitSeconds int,
	heartbeatSeconds int) DoorLogic {

	// https://www.calhoun.io/6-tips-for-using-strings-in-go/
	// https://www.ardanlabs.com/blog/2013/06/gos-duration-type-unravelled.html
	doorOpenWaitDuration, _ := time.ParseDuration(strconv.Itoa(doorOpenWaitSeconds) + "s")
	doorLogic := DoorLogic{doorSensor, onOpen, onClose, onForgot, onHeartbeat, doorOpenWaitDuration, heartbeatSeconds, false}
	return doorLogic
}

// Run starts the DoorLogic object listening.  It'll listen forever.
// Modeled after this example: https://github.com/stianeikeland/go-rpio/blob/master/examples/event/event.go
func (ds DoorLogic) Run() {
	// Get initial sensor state and print
	ds.isOpen = ds.doorSensor.IsOpen()
	ds.printState()

	log.Print("Listening for door sensor")

	// Set initial state for the door open message
	currentForgotTime := new(time.Time)
	if ds.isOpen {
		// Since the door was open when the system started, make sure it closes or send a message
		newForgotTime := time.Now().Add(ds.doorOpenWaitDuration)
		currentForgotTime = &newForgotTime // Need to get reference
	} else {
		currentForgotTime = nil
	}
	isForgotten := false

	// Set up heartbeat
	var nextHeartbeatTime time.Time
	if ds.heartbeatSeconds > 0 {
		nextHeartbeatTime = ds.getNextHeartbeatTime()
		log.Printf("Heartbeat every %v seconds", ds.heartbeatSeconds)
	} else {
		log.Print("Heartbeat disabled")
	}

	// Loop forever listing for changes in the sensor state
	for true {
		// Get new state of sensor
		currentlyOpen := ds.doorSensor.IsOpen()

		// Send forgot alarm if open for too long
		if currentForgotTime != nil && time.Now().After(*currentForgotTime) {
			currentForgotTime = nil
			isForgotten = true
			ds.onForgot()
		}

		// Check if state has changed
		if currentlyOpen != ds.isOpen {
			ds.isOpen = currentlyOpen
			ds.printState()
			if ds.isOpen {
				newForgotTime := time.Now().Add(ds.doorOpenWaitDuration)
				currentForgotTime = &newForgotTime // Need to get reference

				ds.onOpen()
			} else {
				currentForgotTime = nil
				// Only send the door closed message if it was forgotten open
				if isForgotten {
					isForgotten = false
					ds.onClose()
				}
			}
		}

		// Send heartbeat
		if ds.heartbeatSeconds > 0 && time.Now().After(nextHeartbeatTime) {
			ds.onHeartbeat()
			nextHeartbeatTime = ds.getNextHeartbeatTime()
		}

		// Sleep half second before next loop
		time.Sleep(time.Second / 2)
	}

	log.Print("Done listening for door sensor")
}

// getNextHeartbeatTime returns the next time a heartbeat message should be sent
func (ds DoorLogic) getNextHeartbeatTime() time.Time {
	heartbeatDuration, _ := time.ParseDuration(strconv.Itoa(ds.heartbeatSeconds) + "s")
	return time.Now().Add(heartbeatDuration)
}

// printState prints the current open/closed state of the door
func (ds DoorLogic) printState() {
	if ds.isOpen {
		log.Print("Door open")
	} else {
		log.Print("Door closed")
	}
}
