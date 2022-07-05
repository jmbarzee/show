package bender

import (
	"fmt"
	"math"

	"github.com/jmbarzee/show/common"
)

// Sinusoidal is a Bender which provides a single unchanging bend
type Sinusoidal struct {
	Offset    *float64
	Period    *float64
	Amplitude *float64
}

var _ common.Bender = (*Sinusoidal)(nil)

// Bend returns a value representing some change or bend
func (s Sinusoidal) Bend(f float64) float64 {
	cycles := f / *s.Period
	sin := math.Sin(*s.Offset + 2*math.Pi*cycles)
	return *s.Amplitude * sin
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (s *Sinusoidal) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if s.Offset == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.Offset = p.SelectShift()
		})
	}
	if s.Period == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.Period = p.SelectShift()
		})
	}
	if s.Amplitude == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.Amplitude = p.SelectShift()
		})
	}
	return sFuncs
}

func (s Sinusoidal) String() string {
	return fmt.Sprintf("shifter.Sinusoidal{Offset:%v, Period:%v, Amplitude:%v}", s.Offset, s.Period, s.Amplitude)
}
