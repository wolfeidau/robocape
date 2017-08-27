package robocape

/*
#cgo LDFLAGS: -lroboticscape
#include "roboticscape.h"
*/
import "C"

// ESCPulseNormalised send a pulse to the ESC
func ESCPulseNormalised(ch int, throttle float32) error {
	return checkRes(C.rc_send_esc_pulse_normalized(C.int(ch), C.float(throttle)))
}
