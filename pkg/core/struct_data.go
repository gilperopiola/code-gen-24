package core

import (
	"fmt"
	"strings"
)

type StructData struct {
	Name   string
	Fields []StructField
}

type StructField struct {
	Name string
	Type string
}

func (structData StructData) GenerateCode() string {

	var (
		b = strings.Builder{}

		firstLine = "type %s struct {\n"
		fieldLine = "    %s %s\n"
		lastLine  = "}"
	)

	b.WriteString(fmt.Sprintf(firstLine, structData.Name))

	for _, field := range structData.Fields {
		b.WriteString(fmt.Sprintf(fieldLine, field.Name, field.Type))
	}

	b.WriteString(lastLine)

	return b.String()
}
