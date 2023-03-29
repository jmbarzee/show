package painter

import (
	"fmt"
	"math"
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/repeatable"
)

// Bounce is a Painter which provides produces colors bouncing between ColorStart and ColorEnd,
// starting at p.ColorStart and shifting in the direction specified by Up
type Bounce struct {
	ColorStart color.Color
	ColorEnd   color.Color
	Up         *bool
	Shifter    common.Shifter
}

var _ common.Painter = (*Bounce)(nil)

// Paint returns a color based on t and obj
func (p Bounce) Paint(t time.Time, obj common.Renderable) {
	start := p.ColorStart.HSL()
	end := p.ColorEnd.HSL()
	var newC color.HSL
	if *p.Up {
		if start.H < end.H {
			hDistance := end.H - start.H
			sDistance := start.S - end.S
			lDistance := start.L - end.L
			totalShift := p.Shifter.Shift(t, obj)
			bounces := int(totalShift / hDistance)
			remainingShift := math.Mod(totalShift, hDistance)

			var hShift float64
			if (bounces % 2) == 0 {
				// even number of bounces
				hShift = remainingShift
			} else {
				// odd number of bounces
				hShift = hDistance - remainingShift
			}
			hShiftRatio := (hShift / hDistance)
			sShift := sDistance * hShiftRatio
			lShift := lDistance * hShiftRatio

			newC = start
			newC.ShiftHue(hShift)
			newC.SetSaturation(newC.S + sShift)
			newC.SetLightness(newC.L + lShift)

		} else {
			hDistance := (1 - start.H) + end.H
			sDistance := start.S - end.S
			lDistance := start.L - end.L
			totalShift := p.Shifter.Shift(t, obj)
			bounces := int(totalShift / hDistance)
			remainingShift := math.Mod(totalShift, hDistance)

			var hShift float64
			if (bounces % 2) == 0 {
				// even number of bounces
				hShift = -remainingShift
			} else {
				// odd number of bounces
				hShift = -(hDistance - remainingShift)
			}
			hShiftRatio := (hShift / hDistance)
			sShift := sDistance * hShiftRatio
			lShift := lDistance * hShiftRatio

			newC = start
			newC.ShiftHue(-hShift) // shifting past 0
			newC.SetSaturation(newC.S + sShift)
			newC.SetLightness(newC.L + lShift)
		}
	} else {
		if start.H > end.H {
			hDistance := start.H - end.H
			sDistance := start.S - end.S
			lDistance := start.L - end.L
			totalShift := p.Shifter.Shift(t, obj)
			bounces := int(totalShift / hDistance)
			remainingShift := math.Mod(totalShift, hDistance)

			var hShift float64
			if (bounces % 2) == 0 {
				// even number of bounces
				hShift = -remainingShift
			} else {
				// odd number of bounces
				hShift = -(hDistance - remainingShift)
			}
			hShiftRatio := (hShift / hDistance)
			sShift := sDistance * hShiftRatio
			lShift := lDistance * hShiftRatio

			newC = start
			newC.ShiftHue(hShift)
			newC.SetSaturation(newC.S + sShift)
			newC.SetLightness(newC.L + lShift)
		} else {
			hDistance := start.H + (1 - end.H)
			sDistance := start.S - end.S
			lDistance := start.L - end.L
			totalShift := p.Shifter.Shift(t, obj)
			bounces := int(totalShift / hDistance)
			remainingShift := math.Mod(totalShift, hDistance)

			var hShift float64
			if (bounces % 2) == 0 {
				// even number of bounces
				hShift = -remainingShift
			} else {
				// odd number of bounces
				hShift = -(hDistance - remainingShift)
			}
			hShiftRatio := (hShift / hDistance)
			sShift := sDistance * hShiftRatio
			lShift := lDistance * hShiftRatio

			newC = start
			newC.ShiftHue(hShift) // shifting past 0
			newC.SetSaturation(newC.S + sShift)
			newC.SetLightness(newC.L + lShift)
		}
	}
	obj.SetColor(newC)
}

// GetStabilizeFuncs returns StabilizeFunc for all remaining unstablaized traits
func (p *Bounce) GetStabilizeFuncs() []func(p common.Palette) {
	sFuncs := []func(p common.Palette){}
	if p.ColorStart == nil {
		sFuncs = append(sFuncs, func(pa common.Palette) {
			p.ColorStart = pa.SelectColor().HSL()
		})
	}
	if p.ColorEnd == nil {
		sFuncs = append(sFuncs, func(pa common.Palette) {
			p.ColorEnd = pa.SelectColor().HSL()
		})
	}
	if p.Up == nil {
		sFuncs = append(sFuncs, func(pa common.Palette) {
			b := repeatable.Chance(pa.Start(), .5)
			p.Up = &b
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

func (p Bounce) String() string {
	return fmt.Sprintf("painter.Bounce{ColorStart:%v, ColorEnd:%v, Up:%v, Shifter:%v}", p.ColorStart, p.ColorEnd, p.Up, p.Shifter)
}
