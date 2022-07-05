package addressable

import (
	"github.com/jmbarzee/color"
	"github.com/jmbarzee/space"
)

// Basic represents a NeoPixel Light in a line
type Light struct {
	Color        color.Color
	Position     int
	NumPositions int
	Location     space.Cartesian
	Orientation  space.Spherical
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

// GetLocation returns the point in space where the Light is
func (l Light) GetLocation() space.Cartesian {
	return l.Location
}

// GetOrientation returns the direction the Light points
func (l Light) GetOrientation() space.Spherical {
	return l.Orientation
}
