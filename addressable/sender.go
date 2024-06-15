package addressable

import (
	"time"

	"github.com/jmbarzee/show/common/color"
)

type Sender interface {
	Send(Instruction) error
}

type Receiver interface {
	Receive(time.Time) *Instruction
}

type Exchanger struct {
	instructions Instructions
}

var _ Sender = (*Exchanger)(nil)

func NewExchanger() Exchanger {
	return Exchanger{NewInstructions()}
}

func (e Exchanger) Send(instruct Instruction) error {
	e.instructions.Add(instruct)
	return nil
}

func (e Exchanger) Receive(t time.Time) *Instruction {
	return e.instructions.Advance(t)
}

func LightsToColors(lights []*Light) []color.Color {
	colors := make([]color.Color, len(lights))
	for i, light := range lights {
		colors[i] = light.GetColor()
	}
	return colors
}
