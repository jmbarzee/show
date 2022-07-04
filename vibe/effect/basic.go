package effect

import (
	"time"

	"github.com/jmbarzee/color"
	"github.com/jmbarzee/show/ifaces"
	"github.com/jmbarzee/show/vibe/span"
)

type BasicEffect struct {
	span.Span
	Rank int
}

func (e BasicEffect) Priotity() int { return e.Rank }

func (e BasicEffect) Render(time.Time, []ifaces.Light) []color.HSL { return nil }
