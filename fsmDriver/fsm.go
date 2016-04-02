package fsmDriver

import (
	
	"../elevatorDriver"
	//"../userInterface"
	"../queueDriver"
	"fmt"
)

type States int
const(
	undefined States = 0
	atFloor = 1
	doorOpen = 2
	driving = 3
)

func ChannelHandler(chButtonPressed chan elevatorDriver.ButtonStatus){
	

	for{ 
		select{
		case order := <- chButtonPressed:
			queueDriver.AddOrder(order)


		}

	}
}

func Fsm(chEvents chan States, chGetFloor chan int){
	chEvents <- undefined
	for{
		state := <- chEvents
		
		switch(state){
		case undefined:
			fmt.Println("undefined")
			elevatorDriver.ElevInit()
			queueDriver.QueueInit()
			EvStartFloor(chGetFloor, chEvents)
			break

		
		case atFloor:
			fmt.Println("atFloor")
			break
		}	
	}
	

}

func EvStartFloor(chGetFloor chan int, chEvents chan States){
	
	currentFloor := <- chGetFloor
	for currentFloor != -1 && currentFloor > 0{
		elevatorDriver.ElevDrive(-1)
	}
	chEvents <- atFloor

}