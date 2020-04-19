package display_test

import (
	"testing"

	"github.com/cebernardi/arduino/pkg/display"
	"github.com/stretchr/testify/assert"
)

func TestSingleDigitEnhanced_Display(t *testing.T) {
	testCases := []struct {
		name                string
		value               int8
		expectedError       error
		expectedBlinkStatus display.BlinkStatus
		expectedDigit       *display.Digit
	}{
		{
			name:          "it returns an error if value is less than 0",
			value:         -1,
			expectedError: display.ValueRangeError,
		},
		{
			name:          "it returns an error if value is less than 10",
			value:         1,
			expectedError: display.ValueRangeError,
		},
		{
			name:          "it returns an error if value is more than 29",
			value:         40,
			expectedError: display.ValueRangeError,
		},
		{
			name:  "it returns correct config if value is 10",
			value: 10,
			expectedDigit: &display.Digit{
				DP: true,
			},
			expectedBlinkStatus: display.SlowBlink,
		},
		{
			name:  "it returns correct config if value is 20",
			value: 20,
			expectedDigit: &display.Digit{
				DP: true,
			},
			expectedBlinkStatus: display.NoBlink,
		},
		{
			name:  "it returns correct config if value is 30",
			value: 30,
			expectedDigit: &display.Digit{
				DP: true,
			},
			expectedBlinkStatus: display.FastBlink,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sd := display.SingleDigitEnhanced{
				Value: tc.value,
			}
			d, err := sd.Display()
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedDigit, d)
			assert.Equal(t, tc.expectedBlinkStatus, sd.Blink)
		})
	}
}

func TestDigitFromValue(t *testing.T) {
	testCases := []struct {
		name          string
		value         int8
		expectedError error
		expectedDigit *display.Digit
	}{
		{
			name:          "it returns an error if value is not 0-9",
			value:         11,
			expectedError: display.DigitValueRangeError,
		},
		{
			name:          "it returns the expected digit if value is 0",
			value:         0,
			expectedDigit: &display.Digit{DP: true},
		},
		{
			name:  "it returns the expected digit if value is 1",
			value: 1,
			expectedDigit: &display.Digit{
				B: true,
				C: true,
			},
		},
		{
			name:  "it returns the expected digit if value is 2",
			value: 2,
			expectedDigit: &display.Digit{
				A: true,
				B: true,
				G: true,
				E: true,
				D: true,
			},
		},
		{
			name:  "it returns the expected digit if value is 3",
			value: 3,
			expectedDigit: &display.Digit{
				A: true,
				B: true,
				G: true,
				C: true,
				D: true,
			},
		},
		{
			name:  "it returns the expected digit if value is 4",
			value: 4,
			expectedDigit: &display.Digit{
				F: true,
				B: true,
				G: true,
				C: true,
			},
		},
		{
			name:  "it returns the expected digit if value is 5",
			value: 5,
			expectedDigit: &display.Digit{
				A: true,
				F: true,
				G: true,
				C: true,
				D: true,
			},
		},
		{
			name:  "it returns the expected digit if value is 6",
			value: 6,
			expectedDigit: &display.Digit{
				A: true,
				F: true,
				G: true,
				C: true,
				D: true,
				E: true,
			},
		},
		{
			name:  "it returns the expected digit if value is 7",
			value: 7,
			expectedDigit: &display.Digit{
				A: true,
				B: true,
				C: true,
			},
		},
		{
			name:  "it returns the expected digit if value is 8",
			value: 8,
			expectedDigit: &display.Digit{
				A: true,
				B: true,
				C: true,
				D: true,
				E: true,
				F: true,
				G: true,
			},
		},
		{
			name:  "it returns the expected digit if value is 9",
			value: 9,
			expectedDigit: &display.Digit{
				A: true,
				B: true,
				C: true,
				D: true,
				F: true,
				G: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d, err := display.DigitFromValue(tc.value)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedDigit, d)
		})
	}
}
