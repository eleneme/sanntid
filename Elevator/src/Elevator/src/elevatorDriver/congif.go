type ElevButtonType int


const N_FLOORS  	= 	4 
const N_BUTTONS		= 	3   

const (
	BUTTON_CALL_UP ElevButtonType = 0
	BUTTON_CALL_DOWN = 1
	BUTTON_COMMAND = 2
)

type ElevMotorDirection int

const (
	DIR_DOWN ElevMotorDirection = -1
	DIR_UP = 1
	DIR_STOP = 0
)
