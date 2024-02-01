package main

import (
	"log"
	"os"

	"github.com/gilperopiola/code-gen-24/pkg"
)

const (
	defaultInputDir  = "in"
	defaultOutputDir = "out"
)

/* Hey there :) */

func main() {

	// Get input & output directories
	inputDir := getConfig("INPUT_DIR", defaultInputDir)
	outputDir := getConfig("OUTPUT_DIR", defaultOutputDir)

	// Output will be written to this file
	outputFilename := outputDir + "/out.go"

	orchestrator := NewOrchestrator(
		pkg.NewStructFileReader(inputDir),
		pkg.NewStructFileWriter(),
		outputFilename,
	)

	// Run program
	log.Println("3... 2... 1...")
	log.Println("Starting code generation process!")

	if err := orchestrator.Run(); err != nil {
		log.Fatalf("): Code generation process failed: %v", err)
	}

	log.Println("Code successfully written to...")
	log.Println(" - " + outputFilename)
	log.Println("<3")
}

// getConfig retrieves the environment variable or returns a default value
func getConfig(envKey, defaultValue string) string {
	if value, exists := os.LookupEnv(envKey); exists {
		return value
	}
	return defaultValue
}

/* Bye there (: */
