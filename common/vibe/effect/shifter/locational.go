package shifter

import (
	"fmt"
	"time"

	"github.com/jmbarzee/show/common"
)

// Locational is a Shifter which provides shifts that relate to changing time, Directionally
type Locational struct {
	XBender common.Bender
	YBender common.Bender
	ZBender common.Bender
}

var _ common.Shifter = (*Locational)(nil)

// Shift returns a value representing some change or shift
func (s Locational) Shift(t time.Time, obj common.Tangible) float64 {
	loc := obj.GetLocation()
	bendX := s.XBender.Bend(loc.X)
	bendY := s.YBender.Bend(loc.Y)
	bendZ := s.ZBender.Bend(loc.Z)
	return bendX + bendY + bendZ
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (s *Locational) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if s.XBender == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.XBender = p.SelectBender()
		})
	} else {
		sFuncs = append(sFuncs, s.XBender.GetStabilizeFuncs()...)
	}
	if s.YBender == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.YBender = p.SelectBender()
		})
	} else {
		sFuncs = append(sFuncs, s.YBender.GetStabilizeFuncs()...)
	}
	if s.ZBender == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.ZBender = p.SelectBender()
		})
	} else {
		sFuncs = append(sFuncs, s.ZBender.GetStabilizeFuncs()...)
	}
	return sFuncs
}

// Copy returns a deep copy of the Shifter
func (s Locational) Copy() common.Shifter {
	return &Locational{
		XBender: common.CopyBender(s.XBender),
		YBender: common.CopyBender(s.YBender),
		ZBender: common.CopyBender(s.ZBender),
	}
}

// String returns a string representation of the Shifter
func (s Locational) String() string {
	return fmt.Sprintf("shifter.Locational{XBender:%v, YBender:%v, ZBender:%v}", s.XBender, s.YBender, s.ZBender)
}
