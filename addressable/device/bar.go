package device

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/addressable"
	"github.com/jmbarzee/show/addressable/node"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/device"
	"github.com/jmbarzee/show/common/space"
)

const (
	npBarLength  = 2
	ledsPerMeter = 60

	ledsPerNPBar = npBarLength * ledsPerMeter
)

// Bar is a strait bar of lights
type Bar struct {
	*node.Line

	sender addressable.Sender
}

var _ common.Device = (*Bar)(nil)

// Register Bar as node for persistence through JSON
func init() {
	device.Register(func() common.Device { return &Bar{} })
}

// NewBar creates a new Bar
func NewBar(id uuid.UUID, sender addressable.Sender, bearings *space.Object, spacing addressable.Spacing, leds int) *Bar {
	return &Bar{
		Line:   node.NewLine(id, bearings, spacing, leds),
		sender: sender,
	}
}

// GetNodes returns all the Nodes which the device holds
func (b Bar) GetNodes() []common.Node {
	return []common.Node{b.Line}
}

// GetNodeInfos returns all the Nodes which the device holds
func (b Bar) GetNodeInfos() []common.NodeInfo {
	return []common.NodeInfo{b.Line}
}

// Render calls render on the underlying line
func (d Bar) DispatchRender(t time.Time) error {
	allLights := d.Line.Render(t)
	allColors := addressable.LightsToColors(allLights)
	instruction := addressable.Instruction{Time: t, Colors: allColors}
	return d.sender.Send(instruction)
}

// GetType returns the type
func (Bar) GetType() string {
	return "npBar"
}

type barJSON struct {
	*node.Line
}

func (d Bar) MarshalJSON() ([]byte, error) {
	temp := &barJSON{}

	temp.Line = d.Line

	partial, err := json.Marshal(temp)
	if err != nil {
		return nil, err
	}

	wrappedInjection, err := json.Marshal(&struct {
		Type string
	}{
		Type: d.GetType(),
	})
	if err != nil {
		return nil, err
	}

	injection := wrappedInjection[1 : len(wrappedInjection)-1]

	full := bytes.Join([][]byte{
		{'{'},
		injection,
		{','},
		partial[1:],
	}, []byte{})

	// fmt.Printf("%s\n", full)
	return full, nil
}

func (d *Bar) UnmarshalJSON(data []byte) error {
	temp := &barJSON{}

	temp.Line = &node.Line{}

	err := json.Unmarshal(data, temp)
	if err != nil {
		return err
	}

	d.Line = temp.Line
	return nil
}
