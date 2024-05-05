package vibe

import (
	"fmt"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/repeatable"
)

// Basic is a vibe which can produce most Effects
type Basic struct {
	effects []common.Effect

	common.Palette
}

var _ common.Vibe = (*Basic)(nil)

// Duplicate creates a copy of a vibe and insures that
// the duplicated vibe will stabilize/materialize differently
func (v *Basic) Duplicate() common.Vibe {
	newVibe := *v
	(&newVibe).NextSeed()
	return &newVibe
}

// Stabilize locks in part of the visual representation of a vibe.
func (v *Basic) Stabilize() common.Vibe {
	newVibe := *v
	sFuncs := newVibe.GetStabilizeFuncs()
	if len(sFuncs) == 0 {
		return &newVibe
	}
	option := repeatable.Option(newVibe.NextSeed(), len(sFuncs))
	sFuncs[option](&newVibe)
	return &newVibe
}

// Materialize locks all remaining unlocked visuals of a vibe
// then returns the resulting effects
func (v *Basic) Materialize() []common.Effect {
	for {
		sFuncs := v.GetStabilizeFuncs()
		if len(sFuncs) == 0 {
			break
		}
		for _, sFunc := range sFuncs {
			sFunc(v)
		}
	}
	return v.effects
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstabilized traits
func (v *Basic) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	for _, e := range v.effects {
		sFuncs = append(sFuncs, e.GetStabilizeFuncs()...)
	}
	if len(v.effects) == 0 {
		sFuncs = append(sFuncs, func(p common.Palette) {
			v.effects = append(v.effects, p.SelectEffect())
		})
	}
	return sFuncs
}

func (v Basic) String() string {
	s := fmt.Sprintf("vibe.Basic{StartTime:%v, EndTime:%v, Effects:[", v.Start(), v.End())
	for i, e := range v.effects {
		if i != 0 {
			s += ", "
		}
		s += fmt.Sprintf("%v", e)
	}
	s += "]}"
	return s
}
