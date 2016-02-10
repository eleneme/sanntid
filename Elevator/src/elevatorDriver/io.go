
 io.h, channels.go, channels.h and driver.go
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

func ioInit() int {
	return int(C.elev_init())
}

func ioSpeed(speed int) {
	C.elev_set_speed(C.int(speed))
}

func ioGetFloorSensor() int {
	return int(C.elev_get_floor_sensor_signal())
}

func ioGetButtonSignal(button elev_button_type_t, floor int) int {
	return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)))
}

func ioGetStopSignal() int {
	return int(C.elev_get_stop_signal())
}

func ioGetObstruction() int {
	return int(C.elev_get_obstruction_signal())
}

func ioSetFloorIndicator(floor int) {
	C.elev_set_floor_indicator(C.int(floor))
}

func ioSetButtonLamp(button elev_button_type_t, floor int, value int) {
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))
}

func ioSetStopLamp(value int) {
	C.elev_set_stop_lamp(C.int(value))
}

func ioElevSetDoorOpenLamp(value int) {
	C.elev_set_door_open_lamp(C.int(value))
}