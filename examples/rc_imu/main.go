package main

import (
	"fmt"
	"log"

	"github.com/wolfeidau/robocape"
)

func main() {
	err := robocape.Initialise()
	if err != nil {
		log.Fatalf("Unable to initialise cape: %v", err)
	}

	defer robocape.Cleanup()

	imu, err := robocape.InitialiseIMU()
	if err != nil {
		log.Fatalf("Unable to initialise imu: %v", err)
	}

	accelData, err := imu.ReadAccel()
	if err != nil {
		log.Fatalf("Unable to read accel: %v", err)
	}

	fmt.Println("accelData", accelData)

	gyroData, err := imu.ReadGyro()
	if err != nil {
		log.Fatalf("Unable to read gyro: %v", err)
	}

	fmt.Println("gyroData", gyroData)

	magData, err := imu.ReadMag()
	if err != nil {
		log.Fatalf("Unable to read magnetometer: %v", err)
	}

	fmt.Println("magData", magData)

	temp, err := imu.ReadTemp()
	if err != nil {
		log.Fatalf("Unable to read temperature: %v", err)
	}

	fmt.Println("temp", temp)
}
