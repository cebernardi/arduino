package main

import (
	"machine"
)

const (
	led    = machine.LED
	button = machine.Pin(2)
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if button.Get() {
			led.Low()
		} else {
			led.High()
		}
		//time.Sleep(time.Millisecond * 10)
	}
}
