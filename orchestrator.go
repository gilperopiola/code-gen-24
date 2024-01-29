package main

import (
	"strings"

	"github.com/gilperopiola/code-gen-24/pkg/core"
	"github.com/gilperopiola/code-gen-24/pkg/generators"
)

type Orchestrator struct {
	FileReader core.FileReaderI
	FileWriter core.FileWriterI

	StructCodeGenerator generators.CodeGenerator
}

func (o *Orchestrator) GenerateCode() error {

	if err := o.FileReader.ParseFolder(); err != nil {
		return err
	}

	var outputCode strings.Builder
	outputCode.WriteString("package main\n\n")

	for _, generableData := range o.FileReader.GetParsedData() {
		o.StructCodeGenerator.SetData(generableData)

		generatedCode, err := o.StructCodeGenerator.GenerateOutput()
		if err != nil {
			return err
		}

		outputCode.WriteString(generatedCode + "\n\n")
	}

	return o.FileWriter.Write(outputCode.String(), outputFilename)
}
