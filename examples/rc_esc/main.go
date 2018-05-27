package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/wolfeidau/robocape"
)

var (
	ch    = kingpin.Arg("channel", "Servo channel to update.").Required().Int()
	limit = kingpin.Arg("limit", "Servo speed.").Required().Float()
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
		escThrottle float32
	)

	escThrottle = 0

	go func() {

		fmt.Println("Starting with throttle at 0")

		time.Sleep(10 * time.Second)

		var i float32

		// ramp up the throttle over 12 seconds
		for i = 0; i < 0.5; i = i + 0.1 {
			escThrottle = i
			fmt.Println("Changing throttle to", escThrottle)

			time.Sleep(2 * time.Second)
		}

		fmt.Println("Changing throttle to zero")

		escThrottle = 0

		time.Sleep(2 * time.Second)

		fmt.Println("Exiting")
		robocape.Exit()

	}()

	for {
		if robocape.IsExiting() {
			break
		}

		robocape.ESCPulseNormalised(*ch, escThrottle)

		// update the pulse at 20hz
		time.Sleep(50 * time.Millisecond)
	}
}
