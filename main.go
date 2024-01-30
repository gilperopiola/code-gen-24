package main

import (
	"log"
	"os"

	"github.com/gilperopiola/code-gen-24/pkg/core"
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

func main() {
	orchestrator := Orchestrator{
		FileReader: core.NewStructFileReader(inputDir),
		FileWriter: core.NewFileWriter(),
	}

	if err := orchestrator.GenerateCode(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("code successfully written to %s", outputFilename)
}
