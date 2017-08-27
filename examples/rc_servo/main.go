package main

import (
	"log"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/wolfeidau/robocape"
)

var (
	ch    = kingpin.Arg("channel", "Servo channel to update.").Required().Int()
	limit = kingpin.Arg("limit", "Sweep limit.").Required().Float32()
)

func main() {
	kingpin.Parse()

	if *limit < 0.1 || *limit > 1.5 {
		log.Fatalf("Invalid position must be between 0.1 and 1.5")
	}

	err := robocape.Initialise()
	if err != nil {
		log.Fatalf("Unable to initialise cape: %v", err)
	}

	defer robocape.Cleanup()

	err = robocape.EnableServoPowerRail()
	if err != nil {
		log.Fatalf("Unable to enable servo power: %v", err)
	}

	var (
		direction  float32
		servoPos   float32
		sweepLimit float32
	)

	servoPos = 0
	sweepLimit = *limit
	direction = 1

	for {
		if robocape.IsExiting() {
			break
		}

		servoPos = servoPos + (direction * sweepLimit / 20)

		// reset pulse width at end of sweep
		if servoPos > sweepLimit {
			servoPos = sweepLimit
			direction = -1
		} else if servoPos < -(sweepLimit) {
			servoPos = -(sweepLimit)
			direction = 1
		}

		robocape.ServoPulseNormalised(*ch, servoPos)

		// update the pulse at 20hz
		time.Sleep(50 * time.Millisecond)
	}
}
