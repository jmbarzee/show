package common

import (
	"time"

	"github.com/jmbarzee/show/common/color"
)

// Vibe is a heavy abstraction correlating to general feelings in music
type Vibe interface {
	Palette

	Stabilizer

	// Duplicate creates a copy of a vibe and insures that
	// the duplicated vibe will stabilize/materialize differently
	Duplicate() Vibe

	// Stabilize locks in part of the visual representation of a vibe.
	Stabilize() Vibe

	// Materialize locks all remaining unlocked visuals of a vibe
	// then returns the resulting effects
	Materialize() []Effect
}

type Palette interface {
	Seeder

	// SelectEffect returns a Effect
	SelectEffect() Effect
	// SelectPainter returns a Painter
	SelectPainter() Painter
	// SelectBender returns a Bender
	SelectBender() Bender
	// SelectShifter returns a Shifter
	SelectShifter() Shifter

	// SelectColor returns a Color
	SelectColor() color.Color
	// SelectDuration returns a Duration
	// Should generally range from 0.1s to 10s
	SelectDuration() time.Duration
	// SelectShift returns a Shift
	// Should generally range from .01 to 1
	SelectShift() float64
}

// Seeder can produce changing seeds for repeatable randomization
type Seeder interface {
	Spanner
	// NextSeed returns an ever changing seed for repeatable randomization
	NextSeed() time.Time
}

// Spanner is any object which has a beginning and end in time
type Spanner interface {

	// Start returns the Start times
	Start() time.Time
	// End returns the End time
	End() time.Time
}

type Stabilizer interface {
	// GetStabilizeFuncs returns a function for all remaining unstabilized traits
	GetStabilizeFuncs() []func(p Palette)
}
