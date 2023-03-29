package node

import (
	"github.com/jmbarzee/show/addressable"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/node"
	"github.com/jmbarzee/show/common/space"
)

// Line is a representation of a strait line of addressable LEDs.
// A Line with Location {0, 0, 0} and Orientation {1, 0, 0, 0}
// streches from the origin, along the X-axis
// with all LEDs pointing parallel with the Y-axis
type Line struct {
	// row provides the implementation of effect.Allocater
	*row

	// Object provides position and orientation data
	*space.Object

	// Basic provides uuid implementation
	node.Basic
}

var _ common.Node = (*Line)(nil)

// NewLine creates a new Line
func NewLine(bearings *space.Object, spacing Spacing, count int) *Line {
	l := &Line{
		Basic:  node.NewBasic(),
		Object: bearings,
		row:    NewRow(spacing, count),
	}

	l.lightsCache = l.BuildLights()
	return l
}

// SetLocation changes the location of the Line as well as rebuild the lightCache.
// Note that Object.Move() will not rebuild the cache. Use this wrapper instead
func (l *Line) SetLocation(newLocation space.Vector) {
	l.SetLocation(newLocation)
	l.lightsCache = l.BuildLights()
}

// SetOrientation changes the Orientation of the Line as well as rebuild the lightCache.
// Note that Object.Move() will not rebuild the cache. Use this wrapper instead
func (l *Line) SetOrientation(newOrientation space.Quaternion) {
	l.SetOrientation(newOrientation)
	l.lightsCache = l.BuildLights()
}

// Move will change the location and orientation of the Line as well as rebuild the lightCache.
// Note that Object.Move() will not rebuild the cache. Use this wrapper instead
func (l *Line) Move(location space.Vector, orientation space.Quaternion) {
	l.Move(location, orientation)
	l.lightsCache = l.BuildLights()
}

func (l *Line) BuildLights() []*addressable.Light {
	lights := make([]*addressable.Light, l.total)

	thisLEDLocation := l.GetLocation()
	orientation := l.GetOrientation()

	// build relativeLEDTranslation
	relativeLEDTranslation := lineGetNextLEDLocation()
	relativeLEDTranslation.Scale(float64(l.spacing))
	relativeLEDTranslation.Rotate(l.GetOrientation())

	for i := range lights {
		lights[i] = &addressable.Light{
			Position:     i,
			NumPositions: l.total,
			Object: space.NewObject(
				thisLEDLocation,
				orientation,
			),
		}
		thisLEDLocation.Translate(*relativeLEDTranslation)
	}
	return lights
}

// GetType returns the type
func (Line) GetType() string {
	return "NPLine"
}

// lineGetNextLEDLocation
func lineGetNextLEDLocation() *space.Vector {
	return &space.Vector{1, 0, 0}
}
