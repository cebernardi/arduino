package main

import (
	"machine"
	"time"

	"github.com/cebernardi/arduino/pkg/pin"

	"github.com/cebernardi/arduino/pkg/display"
)

const (
	pinA  = machine.Pin(7)
	pinB  = machine.Pin(6)
	pinC  = machine.Pin(3)
	pinD  = machine.Pin(4)
	pinE  = machine.Pin(5)
	pinF  = machine.Pin(8)
	pinG  = machine.Pin(9)
	pinDP = machine.Pin(2)
)

var outputPins = []machine.Pin{pinA, pinB, pinC, pinD, pinE, pinF, pinG, pinDP}

func main() {

	pin.AsOutputPins(outputPins)

	sd := display.SingleDigitEnhanced{
		Value: 50,
	}

	d, b, err := sd.Display()

	if err != nil {
		setPins(d)
		return
	}

	var duration time.Duration

	switch *b {
	case display.NoBlink:
		duration = 0
	case display.FastBlink:
		duration = 200 * time.Millisecond
	case display.SlowBlink:
		duration = 500 * time.Millisecond
	}

	var empty display.Digit
	for {
		setPins(d)
		time.Sleep(duration)

		setPins(&empty)
		time.Sleep(duration)
	}
}

func setPins(d *display.Digit) {
	pinA.Set(d.A)
	pinB.Set(d.B)
	pinC.Set(d.C)
	pinD.Set(d.D)
	pinE.Set(d.E)
	pinF.Set(d.F)
	pinG.Set(d.G)
	pinDP.Set(d.DP)
}
