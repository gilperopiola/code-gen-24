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
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("type %s struct {\n", structData.Name))
	for _, field := range structData.Fields {
		builder.WriteString(fmt.Sprintf("    %s %s\n", field.Name, field.Type))
	}
	builder.WriteString("}")
	return builder.String()
}

var (
	ErrEmptyStructData = fmt.Errorf("error generating struct: empty struct data")
)
