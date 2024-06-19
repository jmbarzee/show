package device

import (
	"encoding/json"
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
	*node.Arc

	sender addressable.Sender
}

var _ common.Device = (*Ring)(nil)

// Register Ring as node for persistence through JSON
func init() {
	device.Register(func() common.Device { return Ring{} })
}

// NewRing returns a new Ring
func NewRing(id uuid.UUID, sender addressable.Sender, bearings *space.Object, spacing addressable.Spacing, leds int, radius float64, aspect addressable.Aspect) Ring {
	d := Ring{
		Arc:    node.NewArc(id, bearings, spacing, leds, radius, aspect),
		sender: sender,
	}

	return d
}

// GetNodes returns all the Nodes which the device holds
func (d Ring) GetNodes() []common.Node {
	return []common.Node{d.Arc}
}

// GetNodeInfos returns all the Nodes which the device holds
func (d Ring) GetNodeInfos() []common.NodeInfo {
	return []common.NodeInfo{d.Arc}
}

// DispatchRender calls render on each of the rings and then appends all the lights
func (d Ring) DispatchRender(t time.Time) error {
	allLights := d.Arc.Render(t)
	allColors := addressable.LightsToColors(allLights)
	return d.sender.Send(addressable.Instruction{Time: t, Colors: allColors})
}

// GetType returns the type
func (d Ring) GetType() string {
	return "npRing"
}

func (d Ring) MarshalJSON() ([]byte, error) {
	temp := &struct {
		*node.Arc
		Type string
	}{}

	temp.Arc = d.Arc
	temp.Type = d.GetType()

	return json.Marshal(temp)
}
