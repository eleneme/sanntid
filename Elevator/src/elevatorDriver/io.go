
package elevatorDriver

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
import "C"

type ElevButtonType int

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


func ioInit() bool {
	return int(C.io_init()) != 0
}

func ioSetBit(channel int) {
	C.io_set_bit(C.int(channel))
}

func ioClearBit(channel int) {
	C.io_clear_bit(C.int(channel))
}

func ioWriteAnalog(channel int, value int) {
	C.io_write_analog(C.int(channel), C.int(value))
}

func ioReadBit(channel int) bool {
	return int(C.io_read_bit(C.int(channel))) != 0
}

func ioReadAnalog(channel int) int {
	return int(C.io_read_analog(C.int(channel)))
}