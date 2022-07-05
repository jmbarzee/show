package node

import (
	"github.com/jmbarzee/show/addressable"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/node"
	"github.com/jmbarzee/space"
)

type basicNode struct{ node.Basic }

// Line is a representation of a strait line of neopixels.
type Line struct {
	// Row provides the implementation of effect.Allocater
	*row
	*space.Object
	basicNode
}

var _ common.Node = (*Line)(nil)

// NewLine creates a new Line
func NewLine(
	length int,
	start space.Cartesian,
	orientation space.Spherical,
	rotation space.Spherical,
) *Line {

	l := &Line{
		basicNode: basicNode{node.NewBasic()},
		Object:    space.NewObject(start, orientation, rotation),
		row:       NewRow(length),
	}

	l.refreshLightsCache()
	return l
}

func (l *Line) Move(location space.Cartesian, orientation, rotation space.Spherical) {
	l.Move(location, orientation, rotation)
	l.refreshLightsCache()
}

func (l *Line) refreshLightsCache() {

	singleLEDVector := l.GetOrientation()
	singleLEDVector.R = distPerLED

	lights := make([]*addressable.Light, l.length)
	for i := range lights {
		thisLEDVector := singleLEDVector.Scale(float64(i))
		lightLocation := l.GetLocation().Translate(thisLEDVector).Cartesian()
		lightOrientation := l.GetRotation()
		lights[i] = &addressable.Light{
			Position:     i,
			NumPositions: l.length,
			Location:     lightLocation,
			Orientation:  lightOrientation,
		}
	}
	l.row.lightsCache = lights
}

// GetType returns the type
func (Line) GetType() string {
	return "NPLine"
}
