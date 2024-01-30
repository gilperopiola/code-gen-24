package main

import (
	"strings"

	"github.com/gilperopiola/code-gen-24/pkg/core"
)

type Orchestrator struct {
	FileReader core.FileReaderI
	FileWriter core.FileWriterI
}

func (o *Orchestrator) GenerateCode() error {
	var outputCode strings.Builder

	if err := o.FileReader.ParseFiles(); err != nil {
		return err
	}

	outputCode.WriteString("package main\n\n")

	for _, generableData := range o.FileReader.GetParsedData() {
		outputCode.WriteString(generableData.GenerateCode() + "\n\n")
	}

	return o.FileWriter.Write(outputCode.String(), outputFilename)
}
