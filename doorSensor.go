package main

import (
    "github.com/stianeikeland/go-rpio/v4"
    "log"
    "strconv"
    "time"
)

type DoorSensor struct {
    onOpen               func()
    onClose              func()
    onForgot             func()
    onHeartbeat          func()
    doorOpenWaitDuration time.Duration
    heartbeatSeconds     int

    isOpen bool
}

// pin is the GPIO pin on the board that we are reading from.
// Wire to 3.3 VDC Power and GPIO3 on P1 Pinout (26-pin Header)
// Raspberry Pi Model A pin guide: https://pi4j.com/1.2/pins/model-a-rev2.html
var (
    // Use mcu pin 22, corresponds to GPIO 3 on the pi
    pin = rpio.Pin(22)
)

// NewDoorSensor creates a new DoorSensor object that you Run() and then it'll fire callbacks on events to the door.
func NewDoorSensor(onOpen func(), onClose func(), onForgot func(), onHeartbeat func(), doorOpenWaitSeconds int, heartbeatSeconds int) DoorSensor {
    // https://www.calhoun.io/6-tips-for-using-strings-in-go/
    // https://www.ardanlabs.com/blog/2013/06/gos-duration-type-unravelled.html
    doorOpenWaitDuration, _ := time.ParseDuration(strconv.Itoa(doorOpenWaitSeconds) + "s")
    doorSensor := DoorSensor{onOpen, onClose, onForgot, onHeartbeat, doorOpenWaitDuration, heartbeatSeconds, false}
    return doorSensor
}

// Run starts the DoorSensor object listening.  It'll listen forever.
// Modeled after this example: https://github.com/stianeikeland/go-rpio/blob/master/examples/event/event.go
func (ds DoorSensor) Run() {
    log.Print("Initializing GPIO")

    // Open and map memory to access gpio, check for errors
    if err := rpio.Open(); err != nil {
        log.Fatalf("Error opening GPIO pin: %v", err)
    }
    // Unmap gpio memory when done
    defer rpio.Close()

    // Initialize pin to read
    pin.Input()

    // Get initial sensor state and print
    ds.isOpen = pin.Read() == rpio.Low
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
                newForgotTime := time.Now().Add(ds.doorOpenWaitDuration)
                currentForgotTime = &newForgotTime // Need to get reference

                ds.onOpen()
            } else {
                currentForgotTime = nil
                ds.onClose()
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
func (ds DoorSensor) getNextHeartbeatTime() time.Time {
    heartbeatDuration, _ := time.ParseDuration(strconv.Itoa(ds.heartbeatSeconds) + "s")
    return time.Now().Add(heartbeatDuration)
}

// printState prints the current open/closed state of the door
func (ds DoorSensor) printState() {
    if ds.isOpen {
        log.Print("Door open")
    } else {
        log.Print("Door closed")
    }
}
