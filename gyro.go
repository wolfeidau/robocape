package robocape

/*
#cgo LDFLAGS: -lroboticscape
#include "roboticscape.h"
*/
import "C"

// CalibrateGyro calibrate the gyro
func CalibrateGyro() error {
	return checkRes(C.rc_calibrate_gyro_routine())
}
