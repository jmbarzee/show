package testutil

import (
	"testing"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/repeat"
)

type StabilizerTest struct {
	Stabilizer       common.Stabilizer
	ExpectedVersions []common.Stabilizer
	Palette          TestPalette
}

func RunStabilizerTest(t *testing.T, c StabilizerTest) {
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
	*repeat.Seed
	Bender            common.Bender
	Color             color.HSL
	Duration          time.Duration
	ShifterHue        common.Shifter
	ShifterLightness  common.Shifter
	ShifterSaturation common.Shifter
	Painter           common.Painter
	Effect            common.Effect
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
func (p TestPalette) SelectDuration() time.Duration {
	return p.Duration
}

// SelectShifter returns a Shifter
func (p TestPalette) SelectShifterHue() common.Shifter {
	return p.ShifterHue
}

// SelectShifter returns a Shifter
func (p TestPalette) SelectShifterLightness() common.Shifter {
	return p.ShifterLightness
}

// SelectShifter returns a Shifter
func (p TestPalette) SelectShifterSaturation() common.Shifter {
	return p.ShifterSaturation
}

// SelectPainter returns a Painter
func (p TestPalette) SelectPainter() common.Painter {
	return p.Painter
}

// SelectEffect returns a Effect
func (p TestPalette) SelectEffect() common.Effect {
	return p.Effect
}
