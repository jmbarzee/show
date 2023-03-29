package device

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/addressable"
	"github.com/jmbarzee/show/addressable/node"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/device"
	"github.com/jmbarzee/show/common/space"
)

// Ring is a ring of LEDs
type Ring struct {
	device.Basic

	*node.Arc

	sender addressable.Sender
}

var _ common.Device = (*Ring)(nil)

// NewRing returns a new Ring
func NewRing(id uuid.UUID, sender addressable.Sender, bearings *space.Object, spacing node.Spacing, leds int, radius float64, aspect node.Aspect) Ring {
	d := Ring{
		Basic:  device.NewBasic(id),
		Arc:    node.NewArc(bearings, spacing, leds, radius, aspect),
		sender: sender,
	}

	return d
}

// GetNodes returns all the Nodes which the device holds
func (d Ring) GetNodes() []common.Node {
	return []common.Node{
		d.Arc,
	}
}

// DispatchRender calls render on each of the rings and then appends all the lights
func (d Ring) DispatchRender(t time.Time) error {
	allLights := append(d.Arc.Render(t))
	allColors := addressable.LightsToColors(allLights)
	return d.sender.Send(addressable.Instruction{Time: t, Colors: allColors})
}

// GetType returns the type
func (d Ring) GetType() string {
	return "npRing"
}
