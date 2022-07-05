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
func (s Combo) Bend(f float64) float64 {
	bend := s.A.Bend(f) + s.B.Bend(f)
	return bend
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (s *Combo) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if s.A == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.A = p.SelectBender()
		})
	} else {
		sFuncs = append(sFuncs, s.A.GetStabilizeFuncs()...)
	}
	if s.B == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.B = p.SelectBender()
		})
	} else {
		sFuncs = append(sFuncs, s.B.GetStabilizeFuncs()...)
	}
	return sFuncs
}

func (s Combo) String() string {
	return fmt.Sprintf("shifter.Combo{A:%v, B:%v}", s.A, s.B)
}
