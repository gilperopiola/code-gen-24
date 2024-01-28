package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gilperopiola/code-gen-24/pkg/core"
	"github.com/gilperopiola/code-gen-24/pkg/generators"
)

/* The perfect code should be a balance of:

-	First of all, it should work. It should behave correctly in most cases, even if it's not entirely correct. It should be usable.
- Second, it should work the way it should.
		- Tautological, I know. It means no hacks or workarounds, no corner cases treated differently if there's a logic solution that works for every scenario.
- On third place code should be understandable. And this includes readability, idiomatic, modularized, concise, inviting.

*/

const (
	inputDir       = "in"
	outputDir      = "out"
	outputFilename = outputDir + "/out.go"

	shouldStopOnErr = false
)

func main() {

	structCodeGen := generators.NewStructCodeGenerator("")

	if err := GenerateCode(structCodeGen); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("code successfully written to %s", outputFilename)
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

		if !strings.Contains(fileName, ".json") {
			continue
		}

		generator.SetSource(inputDir + "/" + fileName)

		if err := generator.Read(); err != nil {
			if !shouldStopOnErr {
				fmt.Print(err)
				continue
			}
			return err
		}

		if err := generator.Generate(); err != nil {
			if !shouldStopOnErr {
				fmt.Print(err)
				continue
			}
			return err
		}

		code.WriteString(generator.GetOutput() + "\n\n")

		generator.Clear()
	}

	/* out */
	if err := os.WriteFile(outputFilename, []byte(code.String()), 0644); err != nil {
		return core.ErrWritingOutput(outputFilename, err)
	}

	return nil
}
