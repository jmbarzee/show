package main

import (
	"encoding/json"
	"os"
	"path"

	"github.com/invopop/jsonschema"
	"github.com/jmbarzee/show/common/vibe"
	"github.com/jmbarzee/show/common/vibe/palette"
)

const (
	schemaFolder = "schema"
)

func main() {

	// Check folder exists
	err := os.MkdirAll(schemaFolder, os.ModePerm)
	if err != nil {
		panic(err)
	}

	GenerateAndWriteSchema(&vibe.Basic{}, "vibe.json")
	GenerateAndWriteSchema(&palette.Basic{}, "palette.json")
}

func GenerateAndWriteSchema(target any, fileName string) {
	reflector := jsonschema.Reflector{}
	err := reflector.AddGoComments("github.com/jmbarzee/show", "./")
	if err != nil {
		panic(err)
	}

	schema := reflector.Reflect(target)
	b, err := json.MarshalIndent(schema, "", "\t")
	if err != nil {
		panic(err)
	}

	vibeFileName := path.Join(schemaFolder, fileName)
	os.WriteFile(vibeFileName, b, 0644)
}
