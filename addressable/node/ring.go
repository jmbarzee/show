package node

import (
	"math"

	"github.com/jmbarzee/show/addressable"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/node"
	"github.com/jmbarzee/space"
)

// Ring is a representation of a ring of neopixels.
// Ring implements effect.Device
type Ring struct {
	// row provides the implementation of effect.Allocater
	*row
	object *space.Object
	basicNode

	// Radius is the distance from the Center to any LED
	radius float64
}

var _ common.Node = (*Ring)(nil)

// NewRing creates a new Ring
func NewRing(
	radius float64,
	center space.Cartesian,
	orientation space.Spherical, // Orientation of center of ring, right hand rule
	rotation space.Spherical, // Orientation of First LED, forced orthogolan to orientation
) *Ring {
	length := int(radius * 2 * math.Pi / distPerLED)

	r := &Ring{
		basicNode: basicNode{node.NewBasic()},
		radius:    radius,
		object:    space.NewObject(center, orientation, rotation),
		row:       NewRow(length),
	}
	r.refreshLightsCache()
	return r
}

func (r *Ring) Move(location space.Cartesian, orientation, rotation space.Spherical) {
	r.object.Move(location, orientation, rotation)
	r.refreshLightsCache()
}

func (r Ring) refreshLightsCache() {

	rotationMatrix := r.object.GetRotation().RotationMatrix()
	translationMatrix := r.object.GetLocation().TranslationMatrix()
	transformationMatrix := translationMatrix.Multiply(rotationMatrix)

	radPerLED := distPerLED / r.radius

	lights := make([]*addressable.Light, r.length)
	for i := range lights {

		phi := radPerLED * float64(i)
		sin, cos := math.Sincos(phi)

		// Location of LED if Ring was in YZ-Plane with first LED on the positive Z axis
		localLocation := space.Cartesian{
			X: 0,
			Y: r.radius * sin,
			Z: r.radius * cos,
		}
		worldLocation := localLocation.Transform(transformationMatrix).Cartesian()

		localOrientation := space.NewSpherical(1, math.Pi/4, phi)

		worldOrientation := localOrientation.Transform(rotationMatrix).Spherical()

		lights[i] = &addressable.Light{
			Position:     i,
			NumPositions: r.length,
			Location:     worldLocation,
			Orientation:  worldOrientation,
		}
	}

	r.lightsCache = lights
}

// GetType returns the type
func (Ring) GetType() string {
	return "NPRing"
}
