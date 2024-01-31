package core

import (
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

// GenerateCode returns the generated code for the struct as a string
func (structData StructData) GenerateCode() string {

	var (
		b = strings.Builder{}

		firstLine = "type %s struct {\n"
		fieldLine = "    %s %s\n"
		lastLine  = "}\n"
	)

	b.WriteString(fmt.Sprintf(firstLine, structData.Name))

	for _, field := range structData.Fields {
		b.WriteString(fmt.Sprintf(fieldLine, field.Name, field.Type))
	}

	b.WriteString(lastLine)

	return b.String()
}
