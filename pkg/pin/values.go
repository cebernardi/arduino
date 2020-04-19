package pin

import "github.com/tinygo-org/tinygo/src/machine"

func AsOutputPins(pins []machine.Pin) {
	for _, pin := range pins {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
}
