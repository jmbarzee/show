package addressable

import (
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/space"
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
