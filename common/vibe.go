package common

import (
	"time"

	"github.com/jmbarzee/color"
)

// Vibe is a heavy abstraction correlating to general feelings in music
type Vibe interface {
	Palette

	Stabalizable

	// Duplicate creates a copy of a vibe and insures that
	// the dupliacted vibe will stabalize/materialize differently
	Duplicate() Vibe

	// Stabilize locks in part of the visual representation of a vibe.
	Stabilize() Vibe

	// Materialize locks all remaining unlocked visuals of a vibe
	// then returns the resulting effects
	Materialize() []Effect
}

type Palette interface {
	Span

	// SelectBender returns a Bender
	SelectBender() Bender
	// SelectColor returns a Color
	SelectColor() color.Color
	// SelectDuration returns a Duration
	// Should generally range from 0.1s to 10s
	SelectDuration() *time.Duration
	// SelectShift returns a Shift
	// Should generally range from .01 to 1
	SelectShift() *float64
	// SelectShifter returns a Shifter
	SelectShifter() Shifter
	// SelectPainter returns a Painter
	SelectPainter() Painter
	// SelectEffect returns a Effect
	SelectEffect() Effect
}

// Span is any object which has a beginning and end in time
type Span interface {
	// Start returns the Start time
	Start() time.Time
	// End returns the End time
	End() time.Time
}
type Stabalizable interface {
	// GetStabilizeFuncs returns a function for all remaining unstablaized traits
	GetStabilizeFuncs() []func(p Palette)
}
