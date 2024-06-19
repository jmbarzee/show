package node

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/addressable"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/node"
	"github.com/jmbarzee/show/common/space"
)

// Arc is a representation of a angular Arc of addressable LEDs.
// An arc with Location {0, 0, 0}, Orientation {1, 0, 0, 0}, and Aspect 0
// rotates about the origin, starting on the positive X-axis,
// towards the positive Y-axis, with LEDs facing inward
type Arc struct {
	// row provides the implementation of effect.Allocater
	*row

	// Object provides position and orientation data
	*space.Object

	// Basic provides uuid implementation
	node.Basic

	// radius is the distance from the Center to any LED
	radius float64

	// aspect is the direction which LEDs on an Arc are facing
	aspect addressable.Aspect
}

var _ common.Node = (*Arc)(nil)

// NewArc creates a new Arc
func NewArc(id uuid.UUID, bearings *space.Object, spacing addressable.Spacing, count int, radius float64, aspect addressable.Aspect) *Arc {

	r := &Arc{
		Basic:  node.Basic{ID: id},
		Object: bearings,
		row:    NewRow(spacing, count),
		radius: radius,
		aspect: aspect,
	}
	r.lightsCache = r.BuildLights()
	return r
}

// GetNodeInfo returns the NodeInfo of this Node or a child node,
// if the given ID is a match
func (a Arc) GetNodeInfo(nodeID uuid.UUID) common.NodeInfo {
	if a.GetID() == nodeID {
		return a
	}
	return nil
}

// Note that Object.Move() will not rebuild the cache. Use this wrapper instead
func (a *Arc) SetLocation(newLocation space.Vector) {
	a.SetLocation(newLocation)
	a.lightsCache = a.BuildLights()
}

// SetOrientation changes the Orientation of the Arc as well as rebuild the lightCache.
// Note that Object.Move() will not rebuild the cache. Use this wrapper instead
func (a *Arc) SetOrientation(newOrientation space.Quaternion) {
	a.SetOrientation(newOrientation)
	a.lightsCache = a.BuildLights()
}

// Move will change the location and orientation of the Arc as well as rebuild the lightCache.
// Note that Object.Move() will not rebuild the cache. Use this wrapper instead

func (a *Arc) Move(location space.Vector, orientation space.Quaternion) {
	a.Move(location, orientation)
	a.lightsCache = a.BuildLights()
}

func (a Arc) BuildLights() []*addressable.Light {
	lights := make([]*addressable.Light, a.total)
	radPerLED := float64(a.spacing) / a.radius

	firstLocation := arcGetFirstLEDLocation()
	firstLocation.Scale(a.radius)
	firstOrientation := space.NewRotationQuaternion(float64(a.aspect), arcGetFirstLEDAspectRotationAxis())

	globalTranslation := a.GetLocation()
	globalRotation := a.GetOrientation()

	for i := range lights {
		phi := radPerLED * float64(i)
		relativeRotation := space.NewRotationQuaternion(phi, arcGetFirstLEDPositionalRotationAxis())

		ledLocation := &firstLocation
		ledOrientation := firstOrientation

		// handle led's relative position in the arc
		ledLocation.Rotate(*relativeRotation)
		ledOrientation = relativeRotation.Cross(*ledOrientation)

		// handle arc's relative values in the world
		ledLocation.Translate(globalTranslation)
		ledOrientation = globalRotation.Cross(*ledOrientation)

		lights[i] = &addressable.Light{
			Position:     i,
			NumPositions: a.total,
			Object:       space.NewObject(*ledLocation, *ledOrientation),
		}
	}

	return lights
}

// GetType returns the type
func (Arc) GetType() string {
	return "NPArc"
}

// arcGetFirstLEDLocation
func arcGetFirstLEDLocation() space.Vector {
	return space.Vector{X: 1, Y: 0, Z: 0}
}

// arcGetFirstLEDLocation
func arcGetFirstLEDAspectRotationAxis() space.Vector {
	return space.Vector{X: 0, Y: 0, Z: 1}
}

// arcGetFirstLEDLocation
func arcGetFirstLEDPositionalRotationAxis() space.Vector {
	return space.Vector{X: 0, Y: 1, Z: 0}
}

type arcJSON struct {
	ID          uuid.UUID
	TotalLights int
	Spacing     addressable.Spacing
	Location    space.Vector
	Orientation space.Quaternion
	Radius      float64
	Aspect      addressable.Aspect
}

func (n *Arc) MarshalJSON() ([]byte, error) {
	temp := &arcJSON{}

	temp.ID = n.ID
	temp.TotalLights = n.total
	temp.Spacing = n.spacing
	temp.Location = n.GetLocation()
	temp.Orientation = n.GetOrientation()
	temp.Radius = n.radius
	temp.Aspect = n.aspect

	return json.Marshal(temp)
}

func (n *Arc) UnmarshalJSON(data []byte) error {
	temp := &arcJSON{}

	if err := json.Unmarshal(data, temp); err != nil {
		return err
	}

	n.ID = temp.ID
	n.Object = space.NewObject(temp.Location, temp.Orientation)
	n.row = NewRow(temp.Spacing, temp.TotalLights)

	n.radius = temp.Radius
	n.aspect = temp.Aspect

	return nil
}
