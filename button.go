package robocape

/*
#cgo LDFLAGS: -lroboticscape
#include "roboticscape.h"
typedef void (*closure)(); // workaround see https://github.com/golang/go/issues/19835
void callPausePressed(void);
void callPauseReleased(void);
void callModePressed(void);
void callModeReleased(void);
*/
import "C"

var (
	pausePressed  = func() {}
	pauseReleased = func() {}
	modePressed   = func() {}
	modeReleased  = func() {}
)

//export goPausePressedCallback
func goPausePressedCallback() {
	pausePressed()
}

//export goPauseReleasedCallback
func goPauseReleasedCallback() {
	pauseReleased()
}

//export goModePressedCallback
func goModePressedCallback() {
	modePressed()
}

//export goModeReleasedCallback
func goModeReleasedCallback() {
	modeReleased()
}

// SetPausePressed register function callback for pause pressed
func SetPausePressed(callback func()) error {
	C.rc_set_pause_pressed_func(C.closure(C.callPausePressed))
	pausePressed = callback
	return nil
}

// SetPauseReleased register function callback for pause released
func SetPauseReleased(callback func()) error {
	C.rc_set_pause_released_func(C.closure(C.callPauseReleased))
	pauseReleased = callback
	return nil
}

// SetModePressed register function callback for mode pressed
func SetModePressed(callback func()) error {
	C.rc_set_mode_pressed_func(C.closure(C.callModePressed))
	modePressed = callback
	return nil
}

// SetModeReleased register function callback for mode released
func SetModeReleased(callback func()) error {
	C.rc_set_mode_released_func(C.closure(C.callModeReleased))
	modeReleased = callback
	return nil
}
