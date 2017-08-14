package robocape

/*
#cgo LDFLAGS: -lroboticscape
#include "roboticscape.h"
*/
import "C"
import "sync"

// IMU holds imu data
// TODO: needs LOCKING
type IMU struct {
	mu   sync.Mutex
	data C.rc_imu_data_t
}

// AccelData data read from accelerometer
type AccelData struct {
	X float32
	Y float32
	Z float32
}

// GyroData data read from gyro
type GyroData struct {
	X float32
	Y float32
	Z float32
}

// MagData data read from magnetometer
type MagData struct {
	X float32
	Y float32
	Z float32
}

// InitialiseIMU setup the IMU
func InitialiseIMU() (*IMU, error) {
	i := &IMU{}

	conf := C.rc_default_imu_config()

	conf.enable_magnetometer = C.int(1)

	return i, checkRes(C.rc_initialize_imu(&i.data, conf))
}

// ReadAccel read the accelerometer data
func (i *IMU) ReadAccel() (*AccelData, error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	ad := &AccelData{}

	err := checkRes(C.rc_read_accel_data(&i.data))
	if err != nil {
		return nil, err
	}

	ad.X = float32(i.data.accel[0])
	ad.Y = float32(i.data.accel[1])
	ad.Z = float32(i.data.accel[2])

	//spew.Dump(i.data)

	return ad, nil
}

// ReadGyro read the gryro data
func (i *IMU) ReadGyro() (*GyroData, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	ad := &GyroData{}

	err := checkRes(C.rc_read_gyro_data(&i.data))
	if err != nil {
		return nil, err
	}

	ad.X = float32(i.data.gyro[0])
	ad.Y = float32(i.data.gyro[1])
	ad.Z = float32(i.data.gyro[2])

	//spew.Dump(i.data)

	return ad, nil
}

// ReadMag read the magnetometer data
func (i *IMU) ReadMag() (*MagData, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	ad := &MagData{}

	err := checkRes(C.rc_read_mag_data(&i.data))
	if err != nil {
		return nil, err
	}

	ad.X = float32(i.data.mag[0])
	ad.Y = float32(i.data.mag[1])
	ad.Z = float32(i.data.mag[2])

	//	spew.Dump(i.data)

	return ad, nil
}

// ReadTemp read the tempreture data
func (i *IMU) ReadTemp() (float32, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	err := checkRes(C.rc_read_imu_temp(&i.data))

	//spew.Dump(i.data)

	return float32(i.data.temp), err
}
