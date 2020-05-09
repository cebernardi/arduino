package pin

import "machine"

func AsOutputPins(pins []machine.Pin) {
	for _, pin := range pins {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
}
