package main

import (
	"log"
	"time"

	"github.com/wolfeidau/robocape"
)

func onPausePressed() {
	log.Println("Pause Pressed")
}

func onPauseReleased() {
	log.Println("Pause Released")
}

func onModePressed() {
	log.Println("Mode Pressed")
}

func onModeReleased() {
	log.Println("Mode Released")
}

func main() {
	err := robocape.Initialise()

	if err != nil {
		log.Fatalf("Unable to initialise cape: %v", err)
	}

	robocape.SetPausePressed(onPausePressed)
	robocape.SetPauseReleased(onPauseReleased)
	robocape.SetModePressed(onModePressed)
	robocape.SetModeReleased(onModeReleased)

	time.Sleep(10 * time.Second)
}
