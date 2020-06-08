package main

import (
	"machine"
	"time"
)

const (
	volts   = 5000.0
	max     = 65535
	samples = 20
)

func main() {
	machine.InitADC()

	pin := machine.ADC0
	pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	adc := machine.ADC{Pin: pin}
	adc.Configure()

	for {
		time.Sleep(250 * time.Millisecond)
		adc.Get()
		var cumulative uint16

		for i := 0; i < 10; i++ {
			time.Sleep(250 * time.Millisecond)
			cumulative += adc.Get()
		}

		avg := cumulative / 10

		x := float32(avg) * (volts / max)

		println(uint(x * 10))

		time.Sleep(3 * time.Second)
	}

}

// stty -F /dev/ttyACM0 9600 -echo
// tinygo flash -target arduino -scheduler tasks cmd/temperature/read-lm35dz.go
