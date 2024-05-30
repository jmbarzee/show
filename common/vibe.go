package common

import (
	"time"

	"github.com/jmbarzee/show/common/color"
)

// Vibe is a heavy abstraction correlating to general feelings in music
type Vibe interface {
	Palette
	Stabilizer
	Spanner

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
	// SelectShifterHue returns a Shifter which is aimed to shift hues
	SelectShifterHue() Shifter
	// SelectShifterLightness returns a Shifter which is aimed to shift lightness
	SelectShifterLightness() Shifter
	// SelectShifterSaturation returns a Shifter which is aimed to shift saturation
	SelectShifterSaturation() Shifter
	// SelectBender returns a Bender
	SelectBender() Bender

	// SelectColor returns a Color
	SelectColor() color.Color
	// SelectDuration returns a Duration
	// Should generally range from 0.1s to 10s
	SelectDuration() time.Duration
}

// Seeder can produce changing seeds for repeatable randomization
type Seeder interface {
	// NextSeed returns an ever changing seed for repeatable randomization
	NextSeed() time.Time

	// Option will choose a option from a range
	// where the options is  [0, options)
	Option(options int) int

	// Chance will determine an outcome based on a percentage [0, 1]
	Chance(chance float64) bool

	// RandDuration will return a random duration [min, max]
	RandDuration(min, max time.Duration) time.Duration

	// RandShift provides a random shift [min, max] with granularity of unit
	RandShift(min, max, unit float64) float64
}

// Spanner is any object which has a beginning and end in time
type Spanner interface {
	// Start returns the Start times
	Start() time.Time
	// SetStart sets the Start time
	SetStart(start time.Time)
	// End returns the End time
	End() time.Time
	// SetEnd sets the End time
	SetEnd(end time.Time)
}

type Stabilizer interface {
	// GetStabilizeFuncs returns a function for all remaining unstabilized traits
	GetStabilizeFuncs() []func(p Palette)
}
