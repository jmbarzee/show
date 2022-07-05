package device

import (
	"time"

	"github.com/jmbarzee/color"
	"github.com/jmbarzee/show/addressable"
)

type Sender interface {
	Send(Instruction) error
}

type Instruction struct {
	t      time.Time
	lights []color.Color
}

func lightsToColors(lights []*addressable.Light) []color.Color {
	colors := make([]color.Color, len(lights))
	for i, light := range lights {
		colors[i] = light.GetColor()
	}
	return colors
}
