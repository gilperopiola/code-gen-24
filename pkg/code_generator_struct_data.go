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

// ParseData takes the file contents as input and parses it into the object
func (structData *StructData) ParseData(jsonData []byte) error {
	decoder := json.NewDecoder(strings.NewReader(string(jsonData)))

	_, err := decoder.Token()
	if err != nil {
		return err
	}

	// Go through each JSON line
	for decoder.More() {
		var token json.Token

		// Field Name
		if token, err = decoder.Token(); err != nil {
			return err
		}
		fieldName, ok := token.(string)
		if !ok {
			return fmt.Errorf("expected struct field name as a string")
		}

		// Field Type
		if token, err = decoder.Token(); err != nil {
			return err
		}
		fieldType, ok := token.(string)
		if !ok {
			return fmt.Errorf("expected struct field type as a string")
		}

		structData.Fields = append(structData.Fields, StructField{Name: fieldName, Type: fieldType})
	}

	return nil
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
