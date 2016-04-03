package main

import (
	"./elevatorDriver"
	"./userInterfaceDriver"
	"./queueDriver"
	"./manager"
	//"fmt"
	//"time"
)

var chButtonPressed = make(chan elevatorDriver.Button)
var chGetFloor = make(chan int)


func main() {
	queueDriver.QueueInit()
	elevatorDriver.ElevInit()

	
	go userInterfaceDriver.NewOrder(chButtonPressed)
	go userInterfaceDriver.FloorTracker(chGetFloor)
	go manager.ChannelHandler(chButtonPressed, chGetFloor)
	
	for{
	

	}
}
