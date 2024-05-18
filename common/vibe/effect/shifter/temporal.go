package shifter

import (
	"fmt"
	"time"

	"github.com/jmbarzee/show/common"
)

// Temporal is a Shifter which provides shifts that relate to changing time, Directionally
type Temporal struct {
	Start    *time.Time
	Interval *time.Duration
	Bender   common.Bender
}

var _ common.Shifter = (*Temporal)(nil)

// Shift returns a value representing some change or shift
func (s Temporal) Shift(t time.Time, obj common.Tangible) float64 {
	secondsPast := float64(t.Sub(*s.Start)) / float64(*s.Interval)
	bend := s.Bender.Bend(secondsPast)
	return bend
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (s *Temporal) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if s.Start == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			t := p.NextSeed()
			s.Start = &t
		})
	}
	if s.Interval == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			t := p.SelectDuration()
			s.Interval = &t
		})
	}
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
func (s Temporal) Copy() common.Shifter {
	return &Temporal{
		Start:    common.CopyTime(s.Start),
		Interval: common.CopyDuration(s.Interval),
		Bender:   common.CopyBender(s.Bender),
	}
}

// String returns a string representation of the Shifter
func (s Temporal) String() string {
	var start, interval string

	if s.Start != nil {
		start = fmt.Sprintf("%v", *s.Start)
	} else {
		start = "<nil>"
	}

	if s.Interval != nil {
		interval = fmt.Sprintf("%v", *s.Interval)
	} else {
		interval = "<nil>"
	}

	return fmt.Sprintf("shifter.Temporal{Start:%v, Interval:%v, Bender:%v}", start, interval, s.Bender)
}
