package shifter

import (
	"github.com/invopop/jsonschema"
	"github.com/jmbarzee/show/common/schemautil"
)

// Shifter is merely a convenience struct for building json schemas
// It is/should not be used for any other purpose
type Shifter struct{}

func (s Shifter) JSONSchema() *jsonschema.Schema {
	// Shifter is used by Painters to change small things over time
	type Shifter struct{}
	return schemautil.AnyOf(Shifter{}, []any{
		Combo{},
		Locational{},
		Positional{},
		Static{},
		Temporal{},
	})
}
