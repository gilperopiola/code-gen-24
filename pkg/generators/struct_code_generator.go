package generators

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gilperopiola/code-gen-24/pkg/core"
)

type CodeGenerator interface {
	SetSource(source string)
	Read() error
	Generate() error
	GetOutput() string
	Clear()
}

type StructCodeGenerator struct {
	fileName string
	data     core.StructData
	output   string
}

func NewStructCodeGenerator(fileName string) *StructCodeGenerator {
	return &StructCodeGenerator{fileName: fileName}
}

func (g *StructCodeGenerator) SetSource(source string) {
	g.fileName = source
}

func (g *StructCodeGenerator) Read() error {
	g.data.Name = strings.TrimSuffix(filepath.Base(g.fileName), filepath.Ext(g.fileName))

	jsonFile, err := os.Open(g.fileName)
	if err != nil {
		return core.ErrOpeningJSONFile(g.fileName, err)
	}
	defer jsonFile.Close()

	g.readIntoStructFields(jsonFile)

	return nil
}

func (g *StructCodeGenerator) readIntoStructFields(jsonFile *os.File) {
	scanner := bufio.NewScanner(jsonFile)

	for scanner.Scan() {
		structField, err := g.parseJSONLineIntoStructField(scanner.Text())
		if err != nil {
			if err != core.ErrParsingShortLine {
				fmt.Println(err)
			}
			continue
		}
		g.data.Fields = append(g.data.Fields, structField)
	}

	if scanner.Err() != nil {
		fmt.Println(fmt.Errorf("error scanning JSON file %s: %v", g.fileName, scanner.Err()))
	}
}

func (g *StructCodeGenerator) parseJSONLineIntoStructField(line string) (core.StructField, error) {
	line = strings.Trim(line, " ,{}")

	if len(strings.Trim(line, " ")) < 4 {
		return core.StructField{}, core.ErrParsingShortLine
	}

	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return core.StructField{}, core.ErrParsingJSONLine(line)
	}

	return core.StructField{
		Name: strings.Trim(parts[0], " \""),
		Type: strings.Trim(parts[1], " \""),
	}, nil
}

func (g *StructCodeGenerator) Generate() error {
	if g.data.Name == "" || len(g.data.Fields) == 0 {
		return core.ErrEmptyStructData(g.fileName)
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("type %s struct {\n", g.data.Name))
	for _, field := range g.data.Fields {
		builder.WriteString(fmt.Sprintf("    %s %s\n", field.Name, field.Type))
	}
	builder.WriteString("}")

	g.output = builder.String()

	return nil
}

func (g *StructCodeGenerator) GetOutput() string {
	return g.output
}

func (g *StructCodeGenerator) Clear() {
	g.fileName = ""
	g.data = core.StructData{}
	g.output = ""
}
