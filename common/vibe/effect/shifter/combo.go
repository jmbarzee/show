package shifter

import (
	"fmt"
	"time"

	"github.com/jmbarzee/show/common"
)

// Combo is a Shifter which provides shifts that relate to changing time, Directionally
type Combo struct {
	A common.Shifter
	B common.Shifter
}

var _ common.Shifter = (*Combo)(nil)

// Shift returns a value representing some change or shift
func (s Combo) Shift(t time.Time, obj common.Tangible) float64 {
	shift := s.A.Shift(t, obj) + s.B.Shift(t, obj)
	return shift
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (s *Combo) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if s.A == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.A = p.SelectShifter()
		})
	} else {
		sFuncs = append(sFuncs, s.A.GetStabilizeFuncs()...)
	}
	if s.B == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.B = p.SelectShifter()
		})
	} else {
		sFuncs = append(sFuncs, s.B.GetStabilizeFuncs()...)
	}
	return sFuncs
}

// Copy returns a deep copy of the Shifter
func (s Combo) Copy() common.Shifter {
	return &Combo{
		A: common.CopyShifter(s.A),
		B: common.CopyShifter(s.B),
	}
}

// String returns a string representation of the Shifter
func (s Combo) String() string {
	return fmt.Sprintf("shifter.Combo{A:%v, B:%v}", s.A, s.B)
}
