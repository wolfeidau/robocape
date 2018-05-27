package robocape

/*
#cgo LDFLAGS: -lroboticscape
#include "roboticscape.h"
*/
import "C"

// EnableMotors enable the Motor
func EnableMotors() error {
	return checkRes(C.rc_enable_motors())
}

// DisableMotors disable the Motor
func DisableMotors() error {
	return checkRes(C.rc_disable_motors())
}

// MotorBrake send a speed to a given motor
func MotorBrake(ch int, throttle float32) error {
	err := checkRes(C.rc_set_motor_brake(C.int(ch)))
	if err != nil {
		return err
	}

	return checkRes(C.rc_set_motor(C.int(ch), C.float(throttle)))
}
