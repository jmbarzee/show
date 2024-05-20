package effect

import (
	"time"

	"github.com/invopop/jsonschema"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/schemautil"
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

// Effect is merely a convenience struct for building json schemas
// It is/should not be used for any other purpose
type Effect struct{}

func (e Effect) JSONSchema() *jsonschema.Schema {
	// Effect is a light abstraction representing patterns of colors
	type Effect struct{}
	return schemautil.AnyOf(Effect{}, []any{
		Solid{},
		Future{},
	})
}
