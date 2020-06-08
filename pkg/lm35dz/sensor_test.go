package lm35dz_test

import (
	"machine"
	"testing"

	"github.com/cebernardi/arduino/pkg/lm35dz"
	"github.com/stretchr/testify/assert"
)

func TestReadTemperature(t *testing.T) {

	testCases := []struct {
		name          string
		pin           machine.Pin
		expectedError error
		expectedTemp  uint16
	}{
		{
			name:          "returns an error if pin is not configured",
			expectedError: lm35dz.ErrPinNotConfigured,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sensor := lm35dz.Sensor{
				Pin: tc.pin,
			}

			temp, err := sensor.ReadTemperature()
			assert.Equal(t, tc.expectedTemp, temp)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
