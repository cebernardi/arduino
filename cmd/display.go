package main

import (
	"machine"

	"github.com/cebernardi/arduino/pkg/display"
)

const (
	pinA  = machine.Pin(7)
	pinB  = machine.Pin(6)
	pinC  = machine.Pin(3)
	pinD  = machine.Pin(9)
	pinE  = machine.Pin(5)
	pinF  = machine.Pin(8)
	pinG  = machine.Pin(9)
	pinDP = machine.Pin(2)
)

func main() {
	pinA.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinB.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinC.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinD.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinE.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinF.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinG.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinDP.Configure(machine.PinConfig{Mode: machine.PinOutput})

	sd := display.SingleDigitEnhanced{
		Value: 12,
	}

	d1, _ := sd.Display()

	pinA.Set(true)
	pinB.Set(true)

	c := 1 % 2
	d := 1 / 2
	pinC.Set(c == 0)
	pinD.Set(d == 0)
	pinE.Set(sd.Value == 1)
	pinE.Set(d1.A)
	// pinC.Set(d.C)
	// pinD.Set(d.D)
	// pinE.Set(d.E)
	// pinF.Set(d.F)
	// pinG.Set(d.G)
	// pinDP.Set(d.DP)

}
