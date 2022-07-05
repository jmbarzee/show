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
			e.TimePerLight = pa.SelectDuration()
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

func (e Future) String() string {
	return fmt.Sprintf("effect.Future{StartTime:%v, EndTime:%v, Rank:%v, TimePerLight:%v, Painter:%v}", e.StartTime, e.EndTime, e.Rank, e.TimePerLight, e.Painter)
}
