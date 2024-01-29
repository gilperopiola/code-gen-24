package generators

import (
	"fmt"
	"strings"

	"github.com/gilperopiola/code-gen-24/pkg/core"
)

type CodeGenerator interface {
	SetData(generable core.Generable)
	GenerateOutput() (string, error)
}

type StructCodeGenerator struct {
	data core.StructData
}

func NewStructCodeGenerator() *StructCodeGenerator {
	return &StructCodeGenerator{}
}

func (g *StructCodeGenerator) SetData(generable core.Generable) {
	g.data = generable.(core.StructData)
}

func (g *StructCodeGenerator) GenerateOutput() (string, error) {
	if g.data.Name == "" || len(g.data.Fields) == 0 {
		return "", core.ErrEmptyStructData
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("type %s struct {\n", g.data.Name))
	for _, field := range g.data.Fields {
		builder.WriteString(fmt.Sprintf("    %s %s\n", field.Name, field.Type))
	}
	builder.WriteString("}")

	return builder.String(), nil
}
