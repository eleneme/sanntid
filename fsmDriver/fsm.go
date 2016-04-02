package fsmDriver

import (
	
	"../elevatorDriver"
	//"../userInterface"
	"../queueDriver"
	"fmt"
	"time"
)

type States int
const(
	
	atFloor = 0
	doorOpen = 1
	driving = 2
)


func ChannelHandler(chButtonPressed chan elevatorDriver.ButtonStatus, chGetFloor chan int, chEvents chan States){
	for{ 
		select{
		case order := <- chButtonPressed:
			queueDriver.AddOrder(order)
			break
		case floor := <- chGetFloor:
			fmt.Println("Current floor",floor)
			arrivedAtFloor(chEvents, floor)
			
			break

		case state := <- chEvents:
			fmt.Println(state)

			switch(state){
			
			case atFloor:
				fmt.Println("State: atFloor")
				
				//GetNextOrder()
				GetInternalOrder(chEvents)
				break


			case doorOpen:
				//queueDriver.DeleteOrder(floor) //kan gjøres i openDoor
				openDoor(chEvents) //closer door etter 3 sek og state til atFloor
				break
			
			case driving:
				
				break

			break	
			}
		}

	}
}

func GetInternalOrder(chEvents chan States){
	fmt.Println("getting internal order")
	currentFloor := elevatorDriver.ElevGetFloorSensorSignal()
	fmt.Println("current floor:", currentFloor)

	for floor := 0; floor < elevatorDriver.N_FLOORS; floor++{
				
		if queueDriver.Queue[floor][elevatorDriver.N_BUTTONS-1] == 1{
			if floor > currentFloor{
				//dir = 1
				//elevatorDriver.ElevStatus.Dir = dir
				elevatorDriver.ElevDrive(1)
				chEvents <-driving
			}else if floor < currentFloor{
				//dir = -1
				//elevatorDriver.ElevStatus.Dir = dir
				elevatorDriver.ElevDrive(-1)
				chEvents <-driving
			}else if floor == currentFloor{
				chEvents <- doorOpen
			}
			
					
			}
		}		
}

func openDoor(chEvents chan States){
	elevatorDriver.ElevSetDoorOpenLamp(1)
	time.Sleep(2*time.Second)
	elevatorDriver.ElevSetDoorOpenLamp(0)
	chEvents <- atFloor

}

func arrivedAtFloor(chEvents chan States, floor int){
	elevatorDriver.ElevSetFloorIndicator(floor)
	fmt.Println("f",floor)
	if queueDriver.EmptyQueue() == true{
		elevatorDriver.ElevDrive(0)
		chEvents <- atFloor
	}else{
		for button := elevatorDriver.BUTTON_CALL_UP; button < elevatorDriver.N_BUTTONS; button++{ //
			if queueDriver.Queue[floor][button] == 1{
				elevatorDriver.ElevDrive(0)
				chEvents <-atFloor
			}
		}	
	}
	
}

/*func Fsm(chGetFloor chan int){

	for{
		state := <- chEvents
		
		switch(state){
	
		case atFloor:
			fmt.Println("State: atFloor")
			if queueDriver.EmptyQueue() == true{
				elevatorDriver.ElevDrive(0)
			}
			//GetNextOrder()
			GetInternalOrder()
			break


		case doorOpen:
			//queueDriver.DeleteOrder(floor) //kan gjøres i openDoor
			openDoor() //closer door etter 3 sek og state til atFloor
			break
		
		case driving:
			//arrivedAtDestination() //sender state til door open - legg til ElevStatus last floor
			break



		}	


	}
	

}
*/
/*
func GetNextOrder(){
	currentFloor := elevatorDriver.ElevStatus.LastFloor
	dir :=  elevatorDriver.ElevStatus.Dir
	if queueDriver.EmptyQueue() == false{
		//Checks if someone orders elevator to current floor
		for button := 0; elevatorDriver.BUTTON_CALL_UP < elevatorDriver.N_BUTTONS-1; button++{ //
			if queueDriver.Queue[currentFloor][button] == 1{
				chEvents <-doorOpen
			}
		}

	}else{
		elevatorDriver.ElevDrive(0)
	}
}*/

