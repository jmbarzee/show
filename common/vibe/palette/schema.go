package palette

import (
	"time"

	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/vibe/effect"
	"github.com/jmbarzee/show/common/vibe/effect/bender"
	"github.com/jmbarzee/show/common/vibe/effect/painter"
	"github.com/jmbarzee/show/common/vibe/effect/shifter"
)

func (v Basic) JSONSchemaProperty(prop string) any {
	switch prop {
	case "Effects":
		return []effect.Effect{}
	case "Painters":
		return []painter.Painter{}
	case "Benders":
		return []bender.Bender{}
	case "Shifters":
		return []shifter.Shifter{}
	case "Colors":
		return []color.HSL{}
	case "Durations":
		return []time.Duration{}
	case "Shifts":
		return []float64{}
	}
	return nil
}
