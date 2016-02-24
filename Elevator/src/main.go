package main

import (
	"./elevatorDriver"
)

func main() {
	elevatorDriver.ElevInit()
	elevatorDriver.ElevDrive(1)
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
