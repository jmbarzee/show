package painter

import (
	"fmt"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
)

// Move is a Painter which provides shifting colors starting at colorStart
type Move struct {
	ColorStart color.Color
	HueShifter common.Shifter
}

var _ common.Painter = (*Move)(nil)

// Paint returns a color based on t and obj
func (p Move) Paint(t time.Time, obj common.Renderable) {
	newColor := p.ColorStart.HSL()
	shift := p.HueShifter.Shift(t, obj)
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
	if p.HueShifter == nil {
		sFuncs = append(sFuncs, func(pa common.Palette) {
			p.HueShifter = pa.SelectShifterHue()
		})
	} else {
		sFuncs = append(sFuncs, p.HueShifter.GetStabilizeFuncs()...)
	}
	return sFuncs
}

// Copy returns a deep copy of the Painter
func (p Move) Copy() common.Painter {
	return &Move{
		ColorStart: common.CopyColor(p.ColorStart),
		HueShifter: common.CopyShifter(p.HueShifter),
	}
}

// String returns a string representation of the Painter
func (p Move) String() string {
	return fmt.Sprintf("painter.Move{ColorStart:%v, Shifter:%v}", p.ColorStart, p.HueShifter)
}
