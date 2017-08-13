package main

import (
	"log"
	"time"

	"github.com/wolfeidau/robocape"
)

func main() {
	err := robocape.Initialise()
	if err != nil {
		log.Fatalf("Unable to initialise cape: %v", err)
	}

	defer robocape.Cleanup()

	log.Println("spin spin spin")

	time.Sleep(2 * time.Second)

	err = robocape.CalibrateGyro()
	if err != nil {
		log.Fatalf("Unable to calibrate gyro: %v", err)
	}

}
