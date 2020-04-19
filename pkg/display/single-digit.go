package display

import (
	"errors"
)

var (
	ValueRangeError      = errors.New("single digit enhanced - value must be between 10 and 39")
	DigitValueRangeError = errors.New("digit - value must be between 0 and 9")
)

const (
	NoBlink   = BlinkStatus("no")
	FastBlink = BlinkStatus("fast")
	SlowBlink = BlinkStatus("slow")
)

type BlinkStatus string

// SingleDigitEnhanced represents a single digit display
// that displays numbers between 10 and 39
// following this convention:
// 10 = slow blinking DP
// 11-19 = slow blinking digit
// 20 = fixed DP
// 21-29 = fixed digit
// 30 = fast blinking DP
// 21-29 = fast blinking digit
type SingleDigitEnhanced struct {
	Value int8
	Blink BlinkStatus
}

type Digit struct {
	A  bool
	B  bool
	C  bool
	D  bool
	E  bool
	F  bool
	G  bool
	DP bool
}

func (sd *SingleDigitEnhanced) Display() (*Digit, error) {
	if sd.Value < 10 {
		return nil, ValueRangeError
	}

	if sd.Value > 39 {
		return nil, ValueRangeError
	}

	mod := sd.Value % 10
	d, err := DigitFromValue(mod)

	if err != nil {
		return nil, err
	}

	times := sd.Value / int8(10)
	if times == 1 {
		sd.Blink = SlowBlink
	} else if times == 2 {
		sd.Blink = NoBlink
	} else if times == 3 {
		sd.Blink = FastBlink
	}

	return d, nil
}

func DigitFromValue(value int8) (*Digit, error) {
	if value < 0 || value > 9 {
		return nil, DigitValueRangeError
	}
	var d Digit
	if value == 0 {
		d.DP = true
	}
	if value == 1 {
		d.B = true
		d.C = true
	}
	if value == 2 {
		d.A = true
		d.B = true
		d.G = true
		d.E = true
		d.D = true
	}
	if value == 3 {
		d.A = true
		d.B = true
		d.G = true
		d.C = true
		d.D = true
	}
	if value == 4 {
		d.F = true
		d.B = true
		d.G = true
		d.C = true
	}
	if value == 5 {
		d.A = true
		d.F = true
		d.G = true
		d.C = true
		d.D = true
	}
	if value == 6 {
		d.A = true
		d.F = true
		d.G = true
		d.E = true
		d.C = true
		d.D = true
	}
	if value == 7 {
		d.A = true
		d.B = true
		d.C = true
	}
	if value == 8 {
		d.A = true
		d.B = true
		d.C = true
		d.D = true
		d.E = true
		d.F = true
		d.G = true
	}
	if value == 9 {
		d.A = true
		d.B = true
		d.C = true
		d.D = true
		d.F = true
		d.G = true
	}
	return &d, nil
}
