package fsmDriver

import (
	
	"../elevatorDriver"
	//"../userInterface"
	"../queueDriver"
	"fmt"
)

type States int
const(
	
	atFloor = 0
	doorOpen = 1
	driving = 2
)


var chEvents = make(chan States)

func ChannelHandler(chButtonPressed chan elevatorDriver.ButtonStatus, chGetFloor chan int){
	for{ 
		select{
		case order := <- chButtonPressed:
			queueDriver.AddOrder(order)
			break
		case floor := <- chGetFloor:
			fmt.Println("Current floor",floor)
			chEvents <- atFloor
			queueDriver.DeleteOrder(floor)
			elevatorDriver.ElevSetFloorIndicator(floor)


		}

	}
}

func Fsm( chGetFloor chan int){

	for{
		state := <- chEvents
		
		switch(state){
	
		case atFloor:
			fmt.Println("State: atFloor")
			if queueDriver.EmptyQueue() == true{
				elevatorDriver.ElevDrive(0)
			}
			break

		case driving:
		}	


	}
	

}


