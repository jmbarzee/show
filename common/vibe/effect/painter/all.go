package painter

import (
	"github.com/invopop/jsonschema"
	"github.com/jmbarzee/show/common/schemautil"
)

// Painter is merely a convenience struct for building json schemas
// It is/should not be used for any other purpose
type Painter struct{}

func (p Painter) JSONSchema() *jsonschema.Schema {
	// Painter is used by effects to select colors
	type Painter struct{}
	return schemautil.AnyOf(Painter{}, []any{
		Bounce{},
		Move{},
		Static{},
	})
}
