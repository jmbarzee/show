package bender

import (
	"fmt"

	"github.com/jmbarzee/show/common"
)

// Combo is a Bender which provides a single unchanging bend
type Combo struct {
	A common.Bender
	B common.Bender
}

var _ common.Bender = (*Combo)(nil)

// Bend returns a value representing some change or bend
func (b Combo) Bend(f float64) float64 {
	bend := b.A.Bend(f) + b.B.Bend(f)
	return bend
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (b *Combo) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if b.A == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			b.A = p.SelectBender()
		})
	} else {
		sFuncs = append(sFuncs, b.A.GetStabilizeFuncs()...)
	}
	if b.B == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			b.B = p.SelectBender()
		})
	} else {
		sFuncs = append(sFuncs, b.B.GetStabilizeFuncs()...)
	}
	return sFuncs
}

// Copy returns a deep copy of the Bender
func (b Combo) Copy() common.Bender {
	return &Combo{
		A: common.CopyBender(b.A),
		B: common.CopyBender(b.B),
	}
}

// String returns a string representation of the Bender
func (b Combo) String() string {
	return fmt.Sprintf("shifter.Combo{A:%v, B:%v}", b.A, b.B)
}
