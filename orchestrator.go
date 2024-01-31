package main

import (
	"strings"

	"github.com/gilperopiola/code-gen-24/pkg/core"
)

// Orchestrator is the struct that ties everything up together
type Orchestrator struct {
	FileReader core.FileReader
	FileWriter core.FileWriter
}

// GenerateCode reads the input files, processes them and writes the generated code to the output file
func (o *Orchestrator) GenerateCode() error {
	var outputCode strings.Builder

	if err := o.FileReader.ParseFiles(); err != nil {
		return err
	}

	outputCode.WriteString("package main\n\n")

	for _, dataUsedAsSource := range o.FileReader.GetParsedData() {
		outputCode.WriteString(dataUsedAsSource.GenerateCode() + "\n")
	}

	return o.FileWriter.Write(outputCode.String(), outputFilename)
}
