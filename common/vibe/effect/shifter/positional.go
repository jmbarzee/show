package shifter

import (
	"fmt"
	"time"

	"github.com/jmbarzee/show/common"
)

// Positional is a Shifter which provides shifts that relate to changing time, Directionally
type Positional struct {
	Bender common.Bender
}

var _ common.Shifter = (*Positional)(nil)

// Shift returns a value representing some change or shift
func (s Positional) Shift(t time.Time, obj common.Tangible) float64 {
	pos, numPos := obj.GetPosition()
	return s.Bender.Bend(float64(pos) / float64(numPos))
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (s *Positional) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if s.Bender == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.Bender = p.SelectBender()
		})
	} else {
		sFuncs = append(sFuncs, s.Bender.GetStabilizeFuncs()...)
	}
	return sFuncs
}

// Copy returns a deep copy of the Shifter
func (s Positional) Copy() common.Shifter {
	return &Positional{
		Bender: common.CopyBender(s.Bender),
	}
}

// String returns a string representation of the Shifter
func (s Positional) String() string {
	return fmt.Sprintf("shifter.Positional{Bender:%v}", s.Bender)
}
