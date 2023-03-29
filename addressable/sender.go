package addressable

import (
	"github.com/jmbarzee/show/common/color"
)

type Sender interface {
	Send(Instruction) error
}

func LightsToColors(lights []*Light) []color.Color {
	colors := make([]color.Color, len(lights))
	for i, light := range lights {
		colors[i] = light.GetColor()
	}
	return colors
}

type Exchanger struct {
	Instructions
}

var _ Sender = (*Exchanger)(nil)

func NewExchanger() Exchanger {
	return Exchanger{NewInstructions()}
}

func (e Exchanger) Send(instruct Instruction) error {
	e.Add(instruct)
	return nil
}
