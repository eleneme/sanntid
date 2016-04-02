package main

import (
	"./elevatorDriver"
	"./userInterfaceDriver"
	//"./queueDriver"
	"./fsmDriver"
	//"fmt"
)


	

func main() {
	chButtonPressed := make(chan elevatorDriver.ButtonStatus)
	chGetFloor := make(chan int)
	chEvents := make(chan fsmDriver.States)

	for{
		go fsmDriver.ChannelHandler(chButtonPressed)
		go userInterfaceDriver.NewOrder(chButtonPressed)
		go userInterfaceDriver.FloorTracker(chGetFloor)
		go fsmDriver.Fsm(chEvents, chGetFloor)

		/*for floor := 0; floor < elevatorDriver.N_FLOORS; floor++ {
			for button := 0; button < elevatorDriver.N_BUTTONS; button++{
				fmt.Print(queueDriver.Queue[floor][button])
			}
			fmt.Println()
		}*/
	
	}
}
/*


func elevStopOutOfBounce() {

	floor := elevatorDriver.ElevGetFloorSensorSignal()
	if floor == 3 {
		elevatorDriver.ElevDrive(0)
		elevatorDriver.ElevDrive(-1)
	} else if floor == 0 {
		elevatorDriver.ElevDrive(0)
		elevatorDriver.ElevDrive(1)
	}
}
*/
