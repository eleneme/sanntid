package elevatorDriver




const N_FLOORS  	= 	4 
const N_BUTTONS		= 	3   

type ElevButtonType int
const (
	BUTTON_CALL_UP ElevButtonType = 0
	BUTTON_CALL_DOWN = 1
	BUTTON_INTERNAL = 2
)

type ElevMotorDirection int
const (
	DIR_DOWN ElevMotorDirection = -1
	DIR_UP = 1
	DIR_STOP = 0
)

type ButtonStatus struct{  //shit navn?
	ButtonType ElevButtonType
	Floor int
}
