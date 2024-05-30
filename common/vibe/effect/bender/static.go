package bender

import (
	"fmt"

	"github.com/jmbarzee/show/common"
)

// Static is a Bender which provides a single unchanging bend
type Static struct {
	TheBend float64
}

var _ common.Bender = (*Static)(nil)

// Bend returns a value representing some change or bend
func (b Static) Bend(_ float64) float64 {
	return b.TheBend
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (b *Static) GetStabilizeFuncs() []func(p common.Palette) {
	return []func(p common.Palette){}
}

// Copy returns a deep copy of the Bender
func (b Static) Copy() common.Bender {
	return &Static{
		TheBend: b.TheBend,
	}
}

// String returns a string representation of the Combo
func (b Static) String() string {
	return fmt.Sprintf("shifter.Static{Bend:%v}", b.TheBend)
}
