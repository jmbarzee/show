package painter

import (
	"fmt"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
)

// Static is a Painter which provides unchanging colors
type Static struct {
	Color color.Color
}

var _ common.Painter = (*Static)(nil)

// Paint returns a color based on t
func (p Static) Paint(t time.Time, obj common.Renderable) {
	obj.SetColor(p.Color.HSL())
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (p *Static) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if p.Color == nil {
		sFuncs = append(sFuncs, func(pa common.Palette) {
			p.Color = pa.SelectColor()
		})
	}
	return sFuncs
}

// Copy returns a deep copy of the Painter
func (p Static) Copy() common.Painter {
	return &Static{
		Color: common.CopyColor(p.Color),
	}
}

// String returns a string representation of the Painter
func (p Static) String() string {
	return fmt.Sprintf("painter.Static{Color:%v}", p.Color)
}
