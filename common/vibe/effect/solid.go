package effect

import (
	"fmt"
	"time"

	"github.com/jmbarzee/show/common"
)

// Solid is an Effect which displays all lights as a single color
type Solid struct {
	BasicEffect
	Painter common.Painter
}

var _ common.Effect = (*Solid)(nil)

// Render will alter obj based on its information and alterability
func (e Solid) Render(t time.Time, obj common.Renderable) {
	e.Painter.Paint(t, obj)
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (e *Solid) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
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
func (e Solid) Copy() common.Effect {
	return &Solid{
		BasicEffect: e.BasicEffect,
		Painter:     common.CopyPainter(e.Painter),
	}
}

// String returns a string representation of the Effect
func (e Solid) String() string {
	return fmt.Sprintf("effect.Solid{StartTime:%v, EndTime:%v, Rank:%v, Painter:%v}", e.Start(), e.End(), e.Rank, e.Painter)
}
