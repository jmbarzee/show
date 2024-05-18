package effect

import (
	"time"

	"github.com/jmbarzee/show/common"
)

type BasicEffect struct {
	common.Spanner
	Rank int
}

func (e BasicEffect) Priority() int { return e.Rank }

// Render will alter obj based on its information and alterability
func (e BasicEffect) Render(time.Time, common.Renderable) {}

// SetSpan sets the start and end of a span from the provided span
func (e *BasicEffect) SetSpan(span common.Spanner) { e.Spanner = span }
