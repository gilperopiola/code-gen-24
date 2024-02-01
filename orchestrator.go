package main

import (
	"fmt"
	"strings"

	"github.com/gilperopiola/code-gen-24/pkg"
)

// Orchestrator is the struct that ties everything up together
type Orchestrator struct {
	FileReader pkg.FileReader
	FileWriter pkg.FileWriter

	OutputFile string
}

// NewOrchestrator creates a new orchestrator instance
func NewOrchestrator(reader pkg.FileReader, writer pkg.FileWriter, output string) *Orchestrator {
	return &Orchestrator{
		FileReader: reader,
		FileWriter: writer,
		OutputFile: output,
	}
}

// Run orchestrates the code generation process
func (o *Orchestrator) Run() error {
	outputCode, err := o.generateCode()
	if err != nil {
		return fmt.Errorf("error generating code: %w", err)
	}

	if err := o.FileWriter.Write(outputCode, o.OutputFile); err != nil {
		return fmt.Errorf("error writing to output file: %w", err)
	}

	return nil
}

// generateCode generates the code string from parsed data
func (o *Orchestrator) generateCode() (string, error) {
	var outputCode strings.Builder
	outputCode.WriteString("package main\n\n")

	if err := o.FileReader.ParseFiles(); err != nil {
		return "", err
	}

	for _, dataUsedAsSource := range o.FileReader.GetParsedData() {
		outputCode.WriteString(dataUsedAsSource.GenerateCode() + "\n")
	}

	return outputCode.String(), nil
}
