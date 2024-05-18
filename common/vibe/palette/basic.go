package palette

import (
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
)

// Basic will randomly choose from the possible options provided
// when the relevant resource is requested from any functions outlined by common.Palette
type Basic struct {
	common.Seeder

	// Effects provides options to SelectEffect
	// Effect spans are not set by the Basic Palette
	Effects []common.Effect
	// Painters provides options to SelectBender
	Painters []common.Painter
	// Benders provides options to SelectBender
	Benders []common.Bender
	// Shifters provides options to SelectShifter
	Shifters []common.Shifter

	// Colors provides options to SelectColor
	Colors []color.Color
	// Durations provides options to SelectDuration
	// Should generally range from 0.1s to 10s
	Durations []time.Duration
	// Shifts provides options to SelectShift
	// Should generally range from .01 to 1
	Shifts []float64
}

var _ common.Palette = (*Basic)(nil)

// SelectEffect returns a Effect
func (v *Basic) SelectEffect() common.Effect {
	options := v.Effects
	length := len(options)
	option := v.Option(length)

	return options[option].Copy()
}

// SelectPainter returns a Painter
func (v *Basic) SelectPainter() common.Painter {
	options := v.Painters
	length := len(options)
	option := v.Option(length)

	return options[option].Copy()
}

// SelectBender returns a Bender
func (v *Basic) SelectBender() common.Bender {
	options := v.Benders
	length := len(options)
	option := v.Option(length)
	return options[option].Copy()
}

// SelectShifter returns a Shifter
func (v *Basic) SelectShifter() common.Shifter {
	options := v.Shifters
	length := len(options)
	option := v.Option(length)

	return options[option].Copy()
}

// SelectColor returns a Color
func (v *Basic) SelectColor() color.Color {
	options := v.Colors
	length := len(options)
	option := v.Option(length)

	return options[option].HSL()
}

// SelectDuration returns a Duration
func (v *Basic) SelectDuration() time.Duration {
	// min := time.Second / 10
	// max := time.Second * 10
	// d := repeat.RandDuration(v.NextSeed(), min, max)
	options := v.Durations
	length := len(options)
	option := v.Option(length)

	return options[option]
}

// SelectShift returns a Shift
func (v *Basic) SelectShift() float64 {
	// min := 0.005
	// max := .5
	// s := repeat.RandShift(v.NextSeed(), min, max, 0.001)
	options := v.Shifts
	length := len(options)
	option := v.Option(length)

	return options[option]
}
