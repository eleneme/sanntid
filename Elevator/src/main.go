package main

import(
	"./elevatorDriver"
	
)

func main(){
	elevatorDriver.ElevInit()

	elevatorDriver.ElevSetButtonLamp(1, 1, 1)
	elevatorDriver.ElevSetMotorDir(-1)
}