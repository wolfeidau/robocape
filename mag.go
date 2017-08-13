package robocape

/*
#cgo LDFLAGS: -lroboticscape
#include "roboticscape.h"
*/
import "C"

// CalibrateMag calibrate the magnetometer
func CalibrateMag() error {
	return checkRes(C.rc_calibrate_mag_routine())
}
