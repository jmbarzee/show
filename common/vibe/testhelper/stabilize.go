package testhelper

import (
	"testing"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/vibe/span"
)

type StabilizeableTest struct {
	Stabilizer       common.Stabilizer
	ExpectedVersions []common.Stabilizer
	Palette          TestPalette
}

func RunStabilizeableTest(t *testing.T, c StabilizeableTest) {
	for i, expectedVersion := range c.ExpectedVersions {

		sFuncs := c.Stabilizer.GetStabilizeFuncs()
		if len(sFuncs) > len(c.ExpectedVersions)-i {
			t.Fatalf("Stabilizer %v failed. Unexpected number of sFuncs:\n\tExpected: %v,\n\tActual: %v", i, len(c.ExpectedVersions)-i, len(sFuncs))
		}
		sFuncs[0](c.Palette)
		if !StructsEqual(c.Stabilizer, expectedVersion) {
			t.Fatalf("Stabilizer %v failed. Stabilizers were not equal:\n\tExpected: %v,\n\tActual: %v", i, expectedVersion, c.Stabilizer)
		}
	}
}

type TestPalette struct {
	span.Span
	Bender   common.Bender
	Color    color.HSL
	Duration time.Duration
	Shift    float64
	Shifter  common.Shifter
	Painter  common.Painter
	Effect   common.Effect
}

// SelectBender returns a Bender
func (p TestPalette) SelectBender() common.Bender {
	return p.Bender
}

// SelectColor returns a Color
func (p TestPalette) SelectColor() color.Color {
	return p.Color
}

// SelectDuration returns a Duration
// Should generally range from 0.1s to 10s
func (p TestPalette) SelectDuration() *time.Duration {
	return &p.Duration
}

// SelectShift returns a Shift
// Should generally range from .01 to 1
func (p TestPalette) SelectShift() *float64 {
	return &p.Shift
}

// SelectShifter returns a Shifter
func (p TestPalette) SelectShifter() common.Shifter {
	return p.Shifter
}

// SelectPainter returns a Painter
func (p TestPalette) SelectPainter() common.Painter {
	return p.Painter
}

// SelectEffect returns a Effect
func (p TestPalette) SelectEffect() common.Effect {
	return p.Effect
}
