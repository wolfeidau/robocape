/*
    This file follows the model  recommended in https://github.com/golang/go/wiki/cgo using a c function 
    to "callback" the associated Go function.
*/
#include "_cgo_export.h"

void
callPausePressed(void)
{
	goPausePressedCallback();
}

void
callPauseReleased(void)
{
	goPauseReleasedCallback();
}

void
callModePressed(void)
{
	goModePressedCallback();
}

void
callModeReleased(void)
{
	goModeReleasedCallback();
}