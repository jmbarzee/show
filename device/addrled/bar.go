package addrled

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/device"
	"github.com/jmbarzee/show/node"
	alednode "github.com/jmbarzee/show/node/addrled"
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

	*alednode.Line
}

var _ device.Device = (*Bar)(nil)

// NewBar creates a new Bar
func NewBar(id uuid.UUID, start space.Cartesian, direction, rotation space.Spherical) Bar {
	return Bar{
		Basic: device.NewBasic(id),
		Line:  alednode.NewLine(ledsPerNPBar, start, direction, rotation),
	}
}

// GetNodes returns all the Nodes which the device holds
func (b Bar) GetNodes() []node.Node {
	return []node.Node{
		b.Line,
	}
}

// Render calls render on the underlying line
func (d Bar) Render(t time.Time) device.Instruction {
	allLights := d.Line.Render(t)
	return instruction{lights: allLights}
}

// GetType returns the type
func (Bar) GetType() string {
	return "npBar"
}
