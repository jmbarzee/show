package bender

import (
	"fmt"
	"math"

	"github.com/jmbarzee/show/common"
)

// Sinusoidal is a Bender which provides a single unchanging bend
type Sinusoidal struct {
	Offset    float64
	Period    float64
	Amplitude float64
}

var _ common.Bender = (*Sinusoidal)(nil)

// Bend returns a value representing some change or bend
func (b Sinusoidal) Bend(f float64) float64 {
	cycles := f / b.Period
	sin := math.Sin(b.Offset + 2*math.Pi*cycles)
	return b.Amplitude * sin
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (b *Sinusoidal) GetStabilizeFuncs() []func(p common.Palette) {
	return []func(p common.Palette){}
}

// Copy returns a deep copy of the Bender
func (b Sinusoidal) Copy() common.Bender {
	return &Sinusoidal{
		Offset:    b.Offset,
		Period:    b.Period,
		Amplitude: b.Amplitude,
	}
}

// String returns a string representation of the Bender
func (b Sinusoidal) String() string {
	return fmt.Sprintf("shifter.Sinusoidal{Offset:%v, Period:%v, Amplitude:%v}", b.Offset, b.Period, b.Amplitude)
}
