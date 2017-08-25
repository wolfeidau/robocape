package robocape

/*
#cgo LDFLAGS: -lroboticscape
#include "roboticscape.h"
typedef void (*closure)(); // workaround see https://github.com/golang/go/issues/19835
void callDmp(void);
*/
import "C"
import "sync"

const radToDegree = 57.295779513

var dmpCallback = func() {}

// IMUDmp holds imu data
// TODO: needs LOCKING
type IMUDmp struct {
	mu   sync.Mutex
	data C.rc_imu_data_t
}

// FusedTaitBryanData data read from DMP
type FusedTaitBryanData struct {
	PitchX float32
	RollY  float32
	YawZ   float32
}

// FusedQuatData data read from DMP
type FusedQuatData struct {
	QuatW float32
	QuatX float32
	QuatY float32
	QuatZ float32
}

// InitialiseIMUDmp setup the IMU with DMP
func InitialiseIMUDmp() (*IMUDmp, error) {
	i := &IMUDmp{}

	conf := C.rc_default_imu_config()

	conf.enable_magnetometer = C.int(1)
	conf.dmp_sample_rate = C.int(50)

	return i, checkRes(C.rc_initialize_imu_dmp(&i.data, conf))
}

// AssignDmpCallback assign a callback function which is invoked for each interupt
func (i *IMUDmp) AssignDmpCallback(callback func(tbd *FusedTaitBryanData, qd *FusedQuatData)) {
	C.rc_set_imu_interrupt_func(C.closure(C.callDmp))
	tbd := &FusedTaitBryanData{}
	qd := &FusedQuatData{}
	dmpCallback = func() {
		tbd.PitchX = float32(i.data.fused_TaitBryan[C.TB_PITCH_X]) * radToDegree
		tbd.RollY = float32(i.data.fused_TaitBryan[C.TB_ROLL_Y]) * radToDegree
		tbd.YawZ = float32(i.data.fused_TaitBryan[C.TB_YAW_Z]) * radToDegree

		qd.QuatW = float32(i.data.fused_quat[C.QUAT_W])
		qd.QuatX = float32(i.data.fused_quat[C.QUAT_X])
		qd.QuatY = float32(i.data.fused_quat[C.QUAT_Y])
		qd.QuatZ = float32(i.data.fused_quat[C.QUAT_Z])
		callback(tbd, qd)
	}
}

//export goDmp
func goDmp() {
	dmpCallback()
}

// rc_initialize_imu_dmp
