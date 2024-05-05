package bender

import (
	"fmt"

	"github.com/jmbarzee/show/common"
)

// Linear is a Bender which provides a single unchanging bend
type Linear struct {
	Interval *float64
}

var _ common.Bender = (*Linear)(nil)

// Bend returns a value representing some change or bend
func (b Linear) Bend(f float64) float64 {
	bend := f / *b.Interval
	return bend
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (b *Linear) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if b.Interval == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			shift := p.SelectShift()
			b.Interval = &shift
		})
	}
	return sFuncs
}

// Copy returns a deep copy of the Bender
func (b Linear) Copy() common.Bender {
	return &Linear{
		Interval: common.CopyFloat64(b.Interval),
	}
}

// String returns a string representation of the Bender
func (b Linear) String() string {
	return fmt.Sprintf("shifter.Linear{Interval:%v}", *b.Interval)
}
