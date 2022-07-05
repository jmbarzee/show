package device

import (
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/addressable/node"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/device"
	"github.com/jmbarzee/space"
)

// ChandelierSmall is a Small Chandelier (2 rings)
type ChandelierSmall struct {
	device.Basic

	*space.Object

	SmallRing *node.Ring
	LargeRing *node.Ring

	sender Sender
}

var _ common.Device = (*ChandelierSmall)(nil)

// NewChandelierSmall returns a new Small Chandelier
func NewChandelierSmall(id uuid.UUID, top space.Cartesian, theta float64) ChandelierSmall {
	orientation := space.NewSpherical(1, theta, 0)
	rotation := space.NewSpherical(1, theta, math.Pi/2)
	d := ChandelierSmall{
		Basic:  device.NewBasic(id),
		Object: space.NewObject(top, orientation, rotation),
	}

	l, o, r := d.getSpaceDataForSmallRing()
	d.SmallRing = node.NewRing(smallRingRadius, l, o, r)

	l, o, r = d.getSpaceDataForLargeRing()
	d.LargeRing = node.NewRing(largeRingRadius, l, o, r)

	return d
}

// GetNodes returns all the Nodes which the device holds
func (d ChandelierSmall) GetNodes() []common.Node {
	return []common.Node{
		d.SmallRing,
		d.LargeRing,
	}
}

// SetLocation changes the physical location of the device
func (d ChandelierSmall) SetLocation(v space.Cartesian) {
	d.Object.SetLocation(v)
	d.updateRingLocations()
}

// SetOrientation changes the physical orientation of the device
// o.Phi will be forced to 0 (because it hangs strait down)
func (d ChandelierSmall) SetOrientation(o space.Spherical) {
	o.P = 0
	d.Object.SetOrientation(o)
	d.updateRingLocations()
}

// SetRotation changes the physical rotation of the device
func (d ChandelierSmall) SetRotation(o space.Spherical) {
	d.Object.SetRotation(o)
	d.updateRingLocations()
}

func (d ChandelierSmall) updateRingLocations() {
	d.SmallRing.Move(d.getSpaceDataForSmallRing())
	d.LargeRing.Move(d.getSpaceDataForLargeRing())
}

func (d ChandelierSmall) getSpaceDataForSmallRing() (space.Cartesian, space.Spherical, space.Spherical) {
	locationTransformation := space.Cartesian{X: 0, Y: 0, Z: smallRingHeight1}.TranslationMatrix()
	location := d.GetLocation().Transform(locationTransformation).Cartesian()

	orientation := d.GetOrientation().Tilt(smallRingTilt1).Rotate(smallRingRotation1)
	rotation := d.GetRotation().Tilt(smallRingTilt1).Rotate(smallRingRotation1)

	return location, orientation, rotation
}

func (d ChandelierSmall) getSpaceDataForLargeRing() (space.Cartesian, space.Spherical, space.Spherical) {
	locationTransformation := space.Cartesian{X: 0, Y: 0, Z: largeRingHeight1}.TranslationMatrix()
	location := d.GetLocation().Transform(locationTransformation).Cartesian()

	orientation := d.GetOrientation().Tilt(largeRingTilt1).Rotate(largeRingRotation1)
	rotation := d.GetRotation().Tilt(largeRingTilt1).Rotate(largeRingRotation1)

	return location, orientation, rotation
}

// DispatchRender calls render on each of the rings and then appends all the lights
func (d ChandelierSmall) DispatchRender(t time.Time) error {
	allLights := append(d.SmallRing.Render(t), d.LargeRing.Render(t)...)
	allColors := lightsToColors(allLights)
	return d.sender.Send(Instruction{t: t, lights: allColors})
}

// GetType returns the type
func (d ChandelierSmall) GetType() string {
	return "npChandelierSmall"
}
