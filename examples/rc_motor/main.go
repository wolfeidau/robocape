package main

import (
	"log"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/wolfeidau/robocape"
)

var (
	ch    = kingpin.Arg("channel", "Motor channel.").Required().Int()
	speed = kingpin.Arg("speed", "Motor speed.").Required().Float32()
)

func main() {
	kingpin.Parse()

	if *speed <= -1.0 || *speed >= 1.0 {
		log.Fatalf("Invalid speed must be between -1.0 and 1.0")
	}

	err := robocape.Initialise()
	if err != nil {
		log.Fatalf("Unable to initialise cape: %v", err)
	}

	err = robocape.EnableMotors()
	if err != nil {
		log.Fatalf("Enable motors: %v", err)
	}

	err = robocape.MotorBrake(*ch, *speed)
	if err != nil {
		log.Fatalf("Unable to run the motor cape: %v", err)
	}

	time.Sleep(5 * time.Second)

	err = robocape.DisableMotors()
	if err != nil {
		log.Fatalf("Enable motors: %v", err)
	}
}
