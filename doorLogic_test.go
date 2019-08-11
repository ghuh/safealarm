package main

import (
	"log"
	"strconv"
	"testing"
	"time"
)

// If the door is open on start up then the forgot to close message should be sent, but not open or closed
func TestInitialOpen(t *testing.T) {
	if !helper(
		TestInitialOpenDoorSensor{},
		true,
		false,
		true,
		false,
		1,
		30,
		1300) {
		t.Error("Failed initial door open test")
	}
}

type TestInitialOpenDoorSensor struct{}

func (ds TestInitialOpenDoorSensor) IsOpen() bool {
	return true
}

// If the door starts closed and stays closed, then there should be no messages if before the hearbeat window
func TestInitialClosed(t *testing.T) {
	if !helper(
		TestInitialClosedDoorSensor{},
		false,
		false,
		false,
		false,
		1,
		30,
		1300) {
		t.Error("Failed initial door closed test")
	}
}

// If the door starts closed and stays closed, then eventually the heartbeat should fire
func TestHeartbeat(t *testing.T) {
	if !helper(
		TestInitialClosedDoorSensor{},
		false,
		false,
		false,
		true,
		1,
		1,
		1300) {
		t.Error("Failed hearbeat")
	}
}

type TestInitialClosedDoorSensor struct{}

func (ds TestInitialClosedDoorSensor) IsOpen() bool {
	return false
}

// Test if the door starts open and then closes, there is no forgot message
func TestDoorOpenThenClosed(t *testing.T) {
	if !helper(
		TestDoorOpenThenClosedSensor{},
		true,
		false,
		false,
		false,
		1,
		30,
		1700) {
		t.Error("Failed open then closed")
	}
}

type TestDoorOpenThenClosedSensor struct{}

var (
	TestDoorOpenThenClosedSensorCount = 0
)

func (ds TestDoorOpenThenClosedSensor) IsOpen() bool {
	TestDoorOpenThenClosedSensorCount++
	// First check it will be open, then the second it will be closed
	if TestDoorOpenThenClosedSensorCount == 1 {
		return true
	}
	return false
}

// Test if the door starts open and then closes eventually, there should be an open, forgot, then closed message
func TestDoorOpenThenSlowClosed(t *testing.T) {
	if !helper(
		TestDoorOpenThenSlowClosedSensor{},
		true,
		true,
		true,
		false,
		1,
		30,
		1700) {
		t.Error("Failed open then slow closed")
	}
}

type TestDoorOpenThenSlowClosedSensor struct{}

var (
	TestDoorOpenThenSlowClosedSensorCount = 0
)

func (ds TestDoorOpenThenSlowClosedSensor) IsOpen() bool {
	TestDoorOpenThenSlowClosedSensorCount++
	// First check it will be open, then the second it will be closed
	if TestDoorOpenThenSlowClosedSensorCount < 4 {
		return true
	}
	return false
}

func helper(
	doorSensor iDoorSensor,
	expectOpen bool,
	expectClosed bool,
	expectForgot bool,
	expectHeartbeat bool,
	doorOpenWaitSeconds int,
	hearbeatSeconds int,
	sleepTimeMillis int) bool {

	sendOpen := false
	sendClosed := false
	sendForgot := false
	sendHeartbeat := false
	doorLogic := NewDoorLogic(
		doorSensor,
		func() { sendOpen = true },
		func() { sendClosed = true },
		func() { sendForgot = true },
		func() { sendHeartbeat = true },
		doorOpenWaitSeconds,
		hearbeatSeconds)
	go doorLogic.Run()

	sleepDuration, _ := time.ParseDuration(strconv.Itoa(sleepTimeMillis) + "ms")
	time.Sleep(sleepDuration)

	if sendOpen != expectOpen {
		log.Printf("Expected 'Open' to be %v, but was %v", expectOpen, sendOpen)
		return false
	}
	if sendClosed != expectClosed {
		log.Printf("Expected 'Closed' to be %v, but was %v", expectClosed, sendClosed)
		return false
	}
	if sendForgot != expectForgot {
		log.Printf("Expected 'Forgot' to be %v, but was %v", expectForgot, sendForgot)
		return false
	}
	if sendHeartbeat != expectHeartbeat {
		log.Printf("Expected 'Heartbeat' to be %v, but was %v", expectHeartbeat, sendHeartbeat)
		return false
	}

	return true
}
