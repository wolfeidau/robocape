package robocape

/*
#cgo LDFLAGS: -lroboticscape
#include "roboticscape.h"
*/
import "C"

// EnableServoPowerRail enable the servo power rail
func EnableServoPowerRail() error {
	return checkRes(C.rc_enable_servo_power_rail())
}

// ServoPulseNormalised send a pulse to the servo
func ServoPulseNormalised(ch int, pos float32) error {
	return checkRes(C.rc_send_servo_pulse_normalized(C.int(ch), C.float(pos)))
}
