package addressable

import (
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/space"
)

// Spacing is an alias for floats to encourage/support common
// LED densities of a strip
type Spacing float64

const (
	// Spacing30 is the spacing of LEDs on a strip which has 30 LEDs per meter
	Spacing30 Spacing = 1.0 / 30.0
	// Spacing60 is the spacing of LEDs on a strip which has 60 LEDs per meter
	Spacing60 Spacing = 1.0 / 60.0
	// Spacing120 is the spacing of LEDs on a strip which has 120 LEDs per meter
	Spacing120 Spacing = 1.0 / 120.0
	// Spacing144 is the spacing of LEDs on a strip which has 144 LEDs per meter
	Spacing144 Spacing = 1.0 / 144.0
)

// Aspect is an alias for float64 to clarify Direction of LEDs in an Arc.
// Given in terms of radians.
type Aspect float64

const (
	// AspectDownward indicates that LEDs on an Arc are facing downward
	AspectDownward Aspect = 0
	// AspectOutward indicates that LEDs on an Arc are facing outward
	AspectOutward Aspect = 0.5
	// AspectUpward indicates that LEDs on an Arc are facing upward
	AspectUpward Aspect = 1.0
	// AspectInward indicates that LEDs on an Arc are facing inward
	AspectInward Aspect = 1.5
)

var _ common.Renderable = (*Light)(nil)

// Basic represents a NeoPixel Light in a line
type Light struct {
	Color        color.Color
	Position     int
	NumPositions int
	*space.Object
}

// GetColor returns the color of the light
func (l Light) GetColor() color.Color {
	return l.Color
}

// SetColor changes the color of the light
func (l *Light) SetColor(newColor color.Color) {
	l.Color = newColor
}

// GetPosition returns the position of the Light (in a string)
func (l Light) GetPosition() (int, int) {
	return l.Position, l.NumPositions
}
