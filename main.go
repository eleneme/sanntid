package main

import (
	"./elevatorDriver"
	"./userInterfaceDriver"
	"./queueDriver"
	"./fsmDriver"
	//"fmt"
	//"time"
)

var chButtonPressed = make(chan elevatorDriver.ButtonStatus)
var chGetFloor = make(chan int)

var chEvents = make(chan fsmDriver.States)

func main() {
	queueDriver.QueueInit()
	elevatorDriver.ElevInit()

	go userInterfaceDriver.NewOrder(chButtonPressed)
	go fsmDriver.ChannelHandler(chButtonPressed, chGetFloor, chEvents)
	go userInterfaceDriver.FloorTracker(chGetFloor)
	//go fsmDriver.Fsm(chGetFloor)
	
	for{
		
		/*for floor := 0; floor < elevatorDriver.N_FLOORS; floor++ {
			for button := 0; button < elevatorDriver.N_BUTTONS; button++{
				fmt.Print(queueDriver.Queue[floor][button])
			}
			fmt.Println()
		}*/
	
	}
}
