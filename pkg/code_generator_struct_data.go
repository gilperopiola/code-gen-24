package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StructData represents the data for a Golang struct
type StructData struct {
	Name   string
	Fields []StructField
}

// StructField represents a field in a Golang struct. Name & Type for now
type StructField struct {
	Name string
	Type string
}

// ParseData takes a JSON file contents as input and parses it into the object
func (structData *StructData) ParseData(jsonData []byte) error {
	decoder := json.NewDecoder(strings.NewReader(string(jsonData)))

	if _, err := decoder.Token(); err != nil {
		return err
	}

	// Go through each JSON line
	for decoder.More() {
		structField, err := parseNextStructField(decoder)
		if err != nil {
			return err
		}

		structData.Fields = append(structData.Fields, structField)
	}

	return nil
}

// parseNextStructField gets 2 tokens from the decoder, the field name and type
func parseNextStructField(decoder *json.Decoder) (StructField, error) {
	var (
		err   error
		out   StructField
		token json.Token
		ok    bool
	)

	// Field Name
	if token, err = decoder.Token(); err != nil {
		return StructField{}, fmt.Errorf("error reading field name: %w", err)
	}
	if out.Name, ok = token.(string); !ok {
		return StructField{}, fmt.Errorf("expected struct field name as a string")
	}

	// Field Type
	if token, err = decoder.Token(); err != nil {
		return StructField{}, fmt.Errorf("error reading field type: %w", err)
	}
	if out.Type, ok = token.(string); !ok {
		return StructField{}, fmt.Errorf("expected struct field type as a string")
	}

	return out, nil
}

// GenerateCode returns the generated code for the struct as a string
func (structData *StructData) GenerateCode() string {
	var b = strings.Builder{}

	b.WriteString(fmt.Sprintf("type %s struct {\n", structData.Name))

	for _, field := range structData.Fields {
		b.WriteString(fmt.Sprintf("    %s %s\n", field.Name, field.Type))
	}

	b.WriteString("}\n")

	return b.String()
}
