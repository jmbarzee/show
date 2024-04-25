package common

import (
	"time"
)

// Effect is a light abstraction representing paterns of colors
type Effect interface {
	Span

	Stabilizer

	// Render will alter obj based on its information and alterability
	// obj is atleast Renderable,
	// but can be any of the other interfaces in renderable.go
	Render(t time.Time, obj Renderable)
	// Priority solves rendering issues
	Priotity() int
}

// Painter is used by effects to select colors
type Painter interface {
	Stabilizer

	// Paint returns a color based on t and obj
	// obj should be atleast Renderable,
	// but can be any of the other interfaces specified in renderable.go
	Paint(t time.Time, obj Renderable)
}

// Shifter is used by Painters to change small things over time
type Shifter interface {
	Stabilizer

	// Shift returns a value representing some change or shift based on t and obj
	// obj should be atleast Tangible,
	// but can be any of the other interfaces specified in renderable.go
	Shift(t time.Time, obj Item) float64
}

// Bender is used by Shifters to change small things over time
type Bender interface {
	Stabilizer

	// Bend returns a value based on f
	Bend(f float64) float64
}
