package bender

import (
	"github.com/invopop/jsonschema"
	"github.com/jmbarzee/show/common/schemautil"
)

// Bender is merely a convenience struct for building json schemas
// It is/should not be used for any other purpose
type Bender struct{}

func (b Bender) JSONSchema() *jsonschema.Schema {
	// Bender is used by Shifters to change small things over time
	type Bender struct{}
	return schemautil.AnyOf(Bender{}, []any{
		Combo{},
		Exponential{},
		Linear{},
		Sinusoidal{},
		Static{},
	})
}
