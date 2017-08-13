package robocape

/*
#cgo LDFLAGS: -lroboticscape
#include "roboticscape.h"
typedef void (*closure)(); // workaround see https://github.com/golang/go/issues/19835
*/
import "C"
import (
	"fmt"
)

// Initialise initialise the robotcape
func Initialise() error {
	return checkRes(C.rc_initialize())
}

// Cleanup release resources used by the robotcape
func Cleanup() error {
	return checkRes(C.rc_cleanup())
}

// checkRes check for non zero return codes
func checkRes(res C.int) error {
	if res != 0 {
		return fmt.Errorf("Call failed res: %v", res)
	}

	return nil
}
