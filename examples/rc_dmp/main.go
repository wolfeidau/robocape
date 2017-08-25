package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/wolfeidau/robocape"
)

var counter = 0

func main() {
	err := robocape.Initialise()
	if err != nil {
		log.Fatalf("Unable to initialise cape: %v", err)
	}

	defer robocape.Cleanup()
	defer robocape.PowerOffIMU()

	imu, err := robocape.InitialiseIMUDmp()
	if err != nil {
		log.Fatalf("Unable to initialise imu: %v", err)
	}

	imu.AssignDmpCallback(printDmp)

	robocape.WaitForExit()
}

func printDmp(tbd *robocape.FusedTaitBryanData, qd *robocape.FusedQuatData) {

	if counter%50 == 0 {
		log.Println("counter:", counter)
		spew.Dump(tbd)
		spew.Dump(qd)
	}
	counter++
}
