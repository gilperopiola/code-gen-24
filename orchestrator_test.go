package main

import (
	"testing"

	"github.com/gilperopiola/code-gen-24/pkg"
)

func TestOrchestrator(t *testing.T) {
	orchestrator := &Orchestrator{
		FileReader: pkg.NewStructFileReader("./in"),
		FileWriter: pkg.NewStructFileWriter(),
		OutputFile: "/out/out.go",
	}

	orchestrator.generateCode()
}
