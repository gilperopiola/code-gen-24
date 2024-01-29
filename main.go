package main

import (
	"log"
	"os"
	"strings"

	"github.com/gilperopiola/code-gen-24/pkg/core"
	"github.com/gilperopiola/code-gen-24/pkg/generators"
)

/* Finish this:

The perfect code should be a balance of:

-	First of all, it should work. It should behave correctly in most cases, even if it's not entirely correct. It should be usable.
- Second, it should work the way it should.
		- Tautological, I know. It means no hacks or workarounds, no corner cases treated differently if there's a logic solution that works for every scenario.
- On third place code should be understandable. And this includes readability, idiomatic, modularized, concise, inviting.
		- And I mean understandable-at-first-sight, or best at second. There will be exceptions, but this should be the rule.


*/

const (
	inputDir       = "in"
	outputDir      = "out"
	outputFilename = outputDir + "/out.go"

	shouldStopOnErr = false
)

type Orchestrator struct {
	StructCodeGenerator generators.CodeGenerator
	FileReader          core.FileReaderI
}

func (o *Orchestrator) GenerateCode() error {
	if err := o.FileReader.ReadInputFolder(); err != nil {
		return err
	}

	for _, structData := range o.FileReader.GetParsedStructData() {
		o.StructCodeGenerator.SetSource(structData.GenerateCode())

		if err := o.StructCodeGenerator.Read(); err != nil {
			return err
		}

		if err := o.StructCodeGenerator.Generate(); err != nil {
			return err
		}

		if err := o.StructCodeGenerator.Write(); err != nil {
			return err
		}

		o.StructCodeGenerator.Clear()
	}

	return nil
}

func main() {
	orchestrator := Orchestrator{
		StructCodeGenerator: generators.NewStructCodeGenerator(""),
		FileReader:          core.NewFileReader(inputDir),
	}

	if err := GenerateCode(structCodeGenerator); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("code successfully written to %s", outputFilename)
}

func GenerateCode(generator generators.CodeGenerator) error {

	/* in */
	inputFiles, err := os.ReadDir(inputDir)
	if err != nil {
		return core.ErrReadingInputFiles(err)
	}

	/* process */
	var code strings.Builder
	code.WriteString("package main\n\n")

	for _, inputFile := range inputFiles {
		fileName := inputFile.Name()

		if !strings.HasSuffix(fileName, ".json") {
			continue
		}

		if err := processFile(fileName, generator, &code); err != nil {
			if shouldStopOnErr {
				return err
			}
			log.Printf("error processing file %s: %v", fileName, err)
		}
	}

	/* out */
	if err := os.WriteFile(outputFilename, []byte(code.String()), 0644); err != nil {
		return core.ErrWritingOutput(outputFilename, err)
	}

	return nil
}

func processFile(fileName string, generator generators.CodeGenerator, code *strings.Builder) error {
	generator.SetSource(inputDir + "/" + fileName)

	if err := generator.Read(); err != nil {
		return err
	}

	if err := generator.Generate(); err != nil {
		return err
	}

	code.WriteString(generator.GetOutput() + "\n\n")

	generator.Clear()

	return nil
}
