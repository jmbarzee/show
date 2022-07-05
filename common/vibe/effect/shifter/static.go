package shifter

import (
	"fmt"
	"time"

	"github.com/jmbarzee/show/common"
)

// Static is a Shifter which provides shifts that relate to changing time, Directionally
type Static struct {
	TheShift *float64
}

var _ common.Shifter = (*Static)(nil)

// Shift returns a value representing some change or shift
func (s Static) Shift(t time.Time, obj common.Item) float64 {
	return *s.TheShift
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (s *Static) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if s.TheShift == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.TheShift = p.SelectShift()
		})
	}
	return sFuncs
}

func (s Static) String() string {
	return fmt.Sprintf("shifter.Static{TheShift:%v}", s.TheShift)
}
