package lm35dz

import (
	"errors"
	"machine"
)

var (
	// ErrBusy is returned when a new read operation is initiated but the unity is still busy
	ErrBusy = errors.New("sensor busy")
	// ErrPinNotConfigured is returned when a read operation is initiated without the pin being configured
	ErrPinNotConfigured = errors.New("pin not configured")
)

// Sensor is the driver for LM35 sensor
type Sensor struct {
	Pin  machine.Pin
	busy bool
}

// ReadTemperature performs 2 analog reads and discards the first one for sake of stability
// return the temperature in millidegrees (celsius)
// it returns ErrBusy if the driver is busy with other measurements
// it returns ConfigErr if Pin is not set
func (s *Sensor) ReadTemperature() (int16, error) {
	return 0, nil
}

// ReadAvgTemperature performs "reads" + 1 analog reads (it discards the first one for sake of stability)
// and returns the average measurements in millidegrees (celsius)
// it returns BusyErr if the driver is busy with other measurements
// it returns ConfigErr if Pin is not set
func (s *Sensor) ReadAvgTemperature(reads int) (int16, error) {
	return 0, nil
}
