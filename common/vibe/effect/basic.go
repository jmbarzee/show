package effect

import (
	"time"

	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/vibe/span"
)

type BasicEffect struct {
	span.Span
	Rank int
}

func (e BasicEffect) Priotity() int { return e.Rank }

// Render will alter obj based on its information and alterability
func (e BasicEffect) Render(time.Time, common.Renderable) {}
