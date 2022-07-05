package painter

import (
	"fmt"
	"time"

	"github.com/jmbarzee/color"
	"github.com/jmbarzee/show/common"
)

// Move is a Painter which provides shifting colors starting at colorStart
type Move struct {
	ColorStart color.Color
	Shifter    common.Shifter
}

var _ common.Painter = (*Move)(nil)

// Paint returns a color based on t and obj
func (p Move) Paint(t time.Time, obj common.Renderable) {
	newColor := p.ColorStart.HSL()
	shift := p.Shifter.Shift(t, obj)
	newColor.ShiftHue(shift)
	obj.SetColor(newColor)
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (p *Move) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if p.ColorStart == nil {
		sFuncs = append(sFuncs, func(pa common.Palette) {
			p.ColorStart = pa.SelectColor()
		})
	}
	if p.Shifter == nil {
		sFuncs = append(sFuncs, func(pa common.Palette) {
			p.Shifter = pa.SelectShifter()
		})
	} else {
		sFuncs = append(sFuncs, p.Shifter.GetStabilizeFuncs()...)
	}
	return sFuncs
}

func (p Move) String() string {
	return fmt.Sprintf("painter.Move{ColorStart:%v, Shifter:%v}", p.ColorStart, p.Shifter)
}
