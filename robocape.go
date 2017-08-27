package robocape

/*
#cgo LDFLAGS: -lroboticscape
#include "roboticscape.h"
typedef void (*closure)(); // workaround see https://github.com/golang/go/issues/19835
*/
import "C"
import (
	"fmt"
	"time"
)

// Initialise initialise the robotcape
func Initialise() error {
	return checkRes(C.rc_initialize())
}

// Cleanup release resources used by the robotcape
func Cleanup() error {
	return checkRes(C.rc_cleanup())
}

// WaitForExit wait for exit
func WaitForExit() {

	for {
		if C.rc_get_state() == C.EXITING {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

}

// IsExiting check if the process should exit
func IsExiting() bool {
	return C.rc_get_state() == C.EXITING
}

// Exit stop the process
func Exit() {
	C.rc_set_state(C.EXITING)
}

// checkRes check for non zero return codes
func checkRes(res C.int) error {
	if res != 0 {
		return fmt.Errorf("Call failed res: %v", res)
	}

	return nil
}
