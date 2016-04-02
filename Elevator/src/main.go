package main

import (
	"./elevatorDriver"
	"fmt"
)

func main() {
	elevatorDriver.ElevInit()
	elevatorDriver.ElevDrive(1)
	fmt.Println("Elevator initialized")
	for {

		elevSetPosition()
		elevStopOutOfBounce()

	}

}

func elevSetPosition() {
	floor := elevatorDriver.ElevGetFloorSensorSignal()
	elevatorDriver.ElevSetFloorIndicator(floor)
	
}

func elevStopOutOfBounce() {

	floor := elevatorDriver.ElevGetFloorSensorSignal()
	if floor == 3 {
		elevatorDriver.ElevDrive(0)
	} else if floor == 0 {
		elevatorDriver.ElevDrive(0)
	}
}
