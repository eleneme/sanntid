package driver  // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.h and driver.go
/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
import "C"

type elev_button_type_t int

const (
	BUTTON_CALL_UP elev_button_type_t = iota
	BUTTON_CALL_DOWN
	BUTTON_COMMAND
)

func Init() int {
	return int(C.elev_init())
}

func Speed(speed int) {
	C.elev_set_speed(C.int(speed))
}

func GetFloorSensor() int {
	return int(C.elev_get_floor_sensor_signal())
}

func GetButtonSignal(button elev_button_type_t, floor int) int {
	return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)))
}

func GetStopSignal() int {
	return int(C.elev_get_stop_signal())
}

func GetObstruction() int {
	return int(C.elev_get_obstruction_signal())
}

func SetFloorIndicator(floor int) {
	C.elev_set_floor_indicator(C.int(floor))
}

func SetButtonLamp(button elev_button_type_t, floor int, value int) {
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))
}

func SetStopLamp(value int) {
	C.elev_set_stop_lamp(C.int(value))
}

func ElevSetDoorOpenLamp(value int) {
	C.elev_set_door_open_lamp(C.int(value))
}