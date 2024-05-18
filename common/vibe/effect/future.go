package effect

import (
	"fmt"
	"time"

	"github.com/jmbarzee/show/common"
)

// Future is an Effect which displays each consecutive light
// as the "future" of the previous light
type Future struct {
	BasicEffect
	TimePerLight *time.Duration
	Painter      common.Painter
}

var _ common.Effect = (*Future)(nil)

// Render will alter obj based on its information and alterability
func (e Future) Render(t time.Time, obj common.Renderable) {
	pos, _ := obj.GetPosition()
	distanceInFuture := *e.TimePerLight * time.Duration(pos)
	e.Painter.Paint(t.Add(distanceInFuture), obj)
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (e *Future) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if e.TimePerLight == nil {
		sFuncs = append(sFuncs, func(pa common.Palette) {
			t := pa.SelectDuration()
			e.TimePerLight = &t
		})
	}
	if e.Painter == nil {
		sFuncs = append(sFuncs, func(pa common.Palette) {
			e.Painter = pa.SelectPainter()
		})
	} else {
		sFuncs = append(sFuncs, e.Painter.GetStabilizeFuncs()...)
	}

	return sFuncs
}

// Copy returns a deep copy of the Effect
func (e Future) Copy() common.Effect {
	return &Future{
		BasicEffect:  e.BasicEffect,
		TimePerLight: common.CopyDuration(e.TimePerLight),
		Painter:      common.CopyPainter(e.Painter),
	}
}

// String returns a string representation of the Effect
func (e Future) String() string {
	var timePerLight string

	if e.TimePerLight != nil {
		timePerLight = fmt.Sprintf("%v", *e.TimePerLight)
	} else {
		timePerLight = "<nil>"
	}
	return fmt.Sprintf("effect.Future{StartTime:%v, EndTime:%v, Rank:%v, TimePerLight:%v, Painter:%v}", e.Start(), e.End(), e.Rank, timePerLight, e.Painter)
}
