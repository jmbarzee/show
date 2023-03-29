package testutil

import (
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/space"
)

// Light represents a Light in a line
type Light struct {
	Color        color.Color
	Position     int
	NumPositions int
	*space.Object
}

var _ common.Renderable = (*Light)(nil)

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

// // GetLocation returns the point in space where the Light is
// func (l Light) GetLocation() space.Vector {
// 	return l.Self.GetLocation()
// }

// // GetOrientation returns the direction the Light points
// func (l Light) GetOrientation() space.Quaternion {
// 	return l.Self.GetOrientation()
// }
