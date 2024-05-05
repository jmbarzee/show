package bender

import (
	"fmt"

	"github.com/jmbarzee/show/common"
)

// Static is a Bender which provides a single unchanging bend
type Static struct {
	TheBend *float64
}

var _ common.Bender = (*Static)(nil)

// Bend returns a value representing some change or bend
func (b Static) Bend(f float64) float64 {
	return *b.TheBend
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (b *Static) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if b.TheBend == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			shift := p.SelectShift()
			b.TheBend = &shift
		})
	}
	return sFuncs
}

// Copy returns a deep copy of the Bender
func (b Static) Copy() common.Bender {
	return &Static{
		TheBend: common.CopyFloat64(b.TheBend),
	}
}

// String returns a string representation of the Combo
func (b Static) String() string {
	return fmt.Sprintf("shifter.Static{Bend:%v}", *b.TheBend)
}
