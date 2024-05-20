package schemautil

import (
	"reflect"
	"strings"

	"github.com/invopop/jsonschema"
)

var reflector jsonschema.Reflector

func init() {
	reflector = jsonschema.Reflector{}
	err := reflector.AddGoComments("github.com/jmbarzee/show", "./")
	if err != nil {
		panic(err)
	}
}

func AnyOf(base any, options []any) *jsonschema.Schema {

	baseSchema := reflector.Reflect(base)

	for _, option := range options {
		schema := reflector.Reflect(option)

		typeName := GetStructName(option)
		typeSchema := &jsonschema.Schema{Const: typeName}
		typeString := "Type"

		// fmt.Printf("%+v\n", schema.Definitions)
		// b, err := json.MarshalIndent(schema.Definitions, "", "\t")
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Printf("%s\n", b)

		actualSchema := schema.Definitions[typeName]
		actualSchema.Properties.Set(typeString, typeSchema)
		actualSchema.Properties.MoveToFront(typeString)
		actualSchema.Required = append(schema.Required, typeString)

		baseSchema.AnyOf = append(baseSchema.AnyOf, actualSchema)
	}

	return baseSchema
}

// GetStructName takes any and returns the name of the struct type
func GetStructName(target any) string {
	t := reflect.TypeOf(target)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return ""
	}
	name := t.String()
	if idx := strings.LastIndex(name, "."); idx != -1 {
		return name[idx+1:]
	}
	return name
}
