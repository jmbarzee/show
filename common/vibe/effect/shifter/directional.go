package shifter

import (
	"fmt"
	"time"

	"github.com/jmbarzee/show/common"
)

// OneShift is just 1
// it can represent a full wrap around of Hue or something else
const OneShift = 1.0

// Directional is a Shifter which provides shifts that relate to changing time, Directionally
type Directional struct {
	PhiBender   common.Bender
	ThetaBender common.Bender
}

var _ common.Shifter = (*Directional)(nil)

// Shift returns a value representing some change or shift
func (s Directional) Shift(t time.Time, obj common.Item) float64 {
	ori := obj.GetOrientation()
	bendPhi := s.PhiBender.Bend(ori.P)
	bendTheta := s.ThetaBender.Bend(ori.T)
	return bendPhi + bendTheta
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (s *Directional) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if s.PhiBender == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.PhiBender = p.SelectBender()
		})
	} else {
		sFuncs = append(sFuncs, s.PhiBender.GetStabilizeFuncs()...)
	}
	if s.ThetaBender == nil {
		sFuncs = append(sFuncs, func(p common.Palette) {
			s.ThetaBender = p.SelectBender()
		})
	} else {
		sFuncs = append(sFuncs, s.ThetaBender.GetStabilizeFuncs()...)
	}
	return sFuncs
}

func (s Directional) String() string {
	return fmt.Sprintf("shifter.Directional{PhiBender:%v, ThetaBender:%v}", s.PhiBender, s.ThetaBender)
}
