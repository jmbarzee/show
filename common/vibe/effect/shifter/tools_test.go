package shifter

import (
	"testing"
	"time"

	"github.com/jmbarzee/show/common"
	helper "github.com/jmbarzee/show/common/vibe/testhelper"
)

type (
	ShiftTest struct {
		Name     string
		Shifter  common.Shifter
		Instants []Instant
	}

	Instant struct {
		Time          time.Time
		Light         common.Tangible
		ExpectedShift float64
	}
)

func RunShifterTests(t *testing.T, cases []ShiftTest) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for i, instant := range c.Instants {
				actualShift := c.Shifter.Shift(instant.Time, instant.Light)
				if !helper.FloatsEqual(instant.ExpectedShift, actualShift, helper.MinErrColor) {
					t.Fatalf("instant %v failed:\n\tExpected: %v,\n\tActual: %v", i, instant.ExpectedShift, actualShift)
				}
			}
		})
	}
}
