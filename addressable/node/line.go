package node

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/addressable"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/node"
	"github.com/jmbarzee/show/common/space"
)

// Line is a representation of a strait line of addressable LEDs.
// A Line with Location {0, 0, 0} and Orientation {1, 0, 0, 0}
// stretches from the origin, along the X-axis
// with all LEDs pointing downwards, parallel with the Y-axis
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
func NewLine(id uuid.UUID, bearings *space.Object, spacing addressable.Spacing, count int) *Line {
	l := &Line{
		Basic:  node.Basic{ID: id},
		Object: bearings,
		row:    NewRow(spacing, count),
	}

	l.lightsCache = l.BuildLights()
	return l
}

// GetNodeInfo returns the NodeInfo of this Node or a child node,
// if the given ID is a match
func (l Line) GetNodeInfo(nodeID uuid.UUID) common.NodeInfo {
	if l.GetID() == nodeID {
		return l
	}
	return nil
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
	return &space.Vector{X: 1, Y: 0, Z: 0}
}

type lineJSON struct {
	ID          uuid.UUID
	TotalLights int
	Spacing     addressable.Spacing
	Location    space.Vector
	Orientation space.Quaternion
}

func (n *Line) MarshalJSON() ([]byte, error) {
	temp := &lineJSON{}

	temp.ID = n.ID
	temp.TotalLights = n.total
	temp.Spacing = n.spacing
	temp.Location = n.GetLocation()
	temp.Orientation = n.GetOrientation()

	return json.Marshal(temp)
}

func (n *Line) UnmarshalJSON(data []byte) error {
	temp := &lineJSON{}

	err := json.Unmarshal(data, temp)
	if err != nil {
		return err
	}

	n.ID = temp.ID
	n.Object = space.NewObject(temp.Location, temp.Orientation)
	n.row = NewRow(temp.Spacing, temp.TotalLights)

	return nil
}
