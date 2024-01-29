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
)

type Orchestrator struct {
	StructCodeGenerator generators.CodeGenerator
	FileReader          core.FileReaderI
}

func main() {
	orchestrator := Orchestrator{
		StructCodeGenerator: generators.NewStructCodeGenerator(),
		FileReader:          core.NewFileReader(inputDir),
	}

	if err := orchestrator.GenerateCode(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("code successfully written to %s", outputFilename)
}

func (o *Orchestrator) GenerateCode() error {

	/* in */
	if err := o.FileReader.ReadInputFolder(); err != nil {
		return err
	}

	var code strings.Builder
	code.WriteString("package main\n\n")

	for _, generableData := range o.FileReader.GetParsedStructData() {

		o.StructCodeGenerator.SetData(generableData)
		generatedCode, err := o.StructCodeGenerator.GenerateOutput()
		if err != nil {
			return err
		}

		code.WriteString(generatedCode + "\n\n")
	}

	/* out */
	if err := os.WriteFile(outputFilename, []byte(code.String()), 0644); err != nil {
		return core.ErrWritingOutput(outputFilename, err)
	}

	return nil
}
