package main

import (
	"machine"
	"time"
)

const flashlightPin = machine.Pin(26) // GPIO pin

func main() {
	flashlightPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		flashlightPin.High()
		time.Sleep(time.Millisecond) // Wait for 1ms
		flashlightPin.Low()
		time.Sleep(5 * time.Second)  // Wait for 5000ms
	}
}
