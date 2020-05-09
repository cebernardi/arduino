package dht11

import (
	"errors"
	"machine"
)

// Sensor represents a HDT11 sensor.
type Sensor interface {
	Read() (*Result, error)
}

// NewSensor builds a new DHT11 sensor
func NewSensor(data machine.Pin) Sensor {
	return &sensor{
		pin: data,
	}
}

// Result contains the sensore readings values
type Result struct {
	Temperature int
	Humidity    int
}

type sensor struct {
	pin  machine.Pin
	busy bool
}

func (s *sensor) Read() (*Result, error) {
	if s.busy {
		return nil, errors.New("busy")
	}

	s.pin.Configure(machine.PinConfig{Mode: machine.PinInput})

	return nil, nil

}
