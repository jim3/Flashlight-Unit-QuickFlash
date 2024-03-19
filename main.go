package main

import (
	"machine"
	"time"
)

const flashlightPin = machine.Pin(26) // GPIO pin number for the flashlight

func main() {
	flashlightPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		flashlightPin.High()         // Turn on the flashlight
		time.Sleep(time.Millisecond) // Wait for 1ms
		flashlightPin.Low()          // Turn off the flashlight
		time.Sleep(5 * time.Second)  // Wait for 5000ms
	}
}
