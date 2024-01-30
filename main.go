package main

import (
	"log"
	"os"

	"github.com/gilperopiola/code-gen-24/pkg/core"
)

const (
	inputDir       = "in"
	outputDir      = "out"
	outputFilename = outputDir + "/out.go"
)

func main() {
	orchestrator := Orchestrator{
		FileReader: core.NewStructFileReader(inputDir),
		FileWriter: core.NewStructFileWriter(),
	}

	if err := orchestrator.GenerateCode(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("code successfully written to %s", outputFilename)
}
