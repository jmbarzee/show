package device

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/addressable/node"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/device"
	"github.com/jmbarzee/space"
)

const (
	npBarLength  = 2
	ledsPerMeter = 60

	ledsPerNPBar = npBarLength * ledsPerMeter
)

// Bar is a strait bar of lights
type Bar struct {
	device.Basic

	*node.Line

	sender Sender
}

var _ common.Device = (*Bar)(nil)

// NewBar creates a new Bar
func NewBar(id uuid.UUID, start space.Cartesian, direction, rotation space.Spherical, sender Sender) Bar {
	return Bar{
		Basic:  device.NewBasic(id),
		Line:   node.NewLine(ledsPerNPBar, start, direction, rotation),
		sender: sender,
	}
}

// GetNodes returns all the Nodes which the device holds
func (b Bar) GetNodes() []common.Node {
	return []common.Node{b.Line}
}

// Render calls render on the underlying line
func (d Bar) DispatchRender(t time.Time) error {
	allLights := d.Line.Render(t)
	allColors := lightsToColors(allLights)
	return d.sender.Send(Instruction{t: t, lights: allColors})
}

// GetType returns the type
func (Bar) GetType() string {
	return "npBar"
}
