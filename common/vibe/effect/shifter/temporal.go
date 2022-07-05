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
func (s Temporal) Shift(t time.Time, obj common.Item) float64 {
	secondsPast := float64(t.Sub(*s.Start)) / float64(*s.Interval)
	bend := s.Bender.Bend(secondsPast)
	return bend
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (s *Temporal) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if s.Start == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			t := p.Start()
			s.Start = &t
		})
	}
	if s.Interval == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.Interval = p.SelectDuration()
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

func (s Temporal) String() string {
	return fmt.Sprintf("shifter.Temporal{Start:%v, Interval:%v, Bender:%v}", s.Start, s.Interval, s.Bender)
}
