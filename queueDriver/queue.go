package queueDriver

import (
	"../elevatorDriver"
	//"fmt"
)

var Queue = [elevatorDriver.N_FLOORS][elevatorDriver.N_BUTTONS]int{}

func QueueInit(){
	Queue = [elevatorDriver.N_FLOORS][elevatorDriver.N_BUTTONS]int{{0, -1, 0}, {0, 0, 0}, {0, 0, 0}, {-1, 0, 0}}
}

func AddOrder(order elevatorDriver.ButtonStatus){
	Queue[order.Floor][order.ButtonType] = 1
	elevatorDriver.ElevSetButtonLamp(order.Floor, order.ButtonType, 1)
}

/*func DeleteOrder(){

}

func StopAtFloor(){

}

func EmptyQueue(){

}

func OrderAbove(){

}

func OrderBelow(){
	
}*/