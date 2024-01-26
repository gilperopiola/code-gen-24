package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

	shouldStopOnError = false
)

func main() {
	if err := GenerateCode(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("code successfully written to %s", outputFilename)
}

func GenerateCode() error {

	inputFiles, err := readInputFiles(inputDir)
	if err != nil {
		return errReadingInputFiles(err)
	}

	var code strings.Builder
	code.WriteString("package main\n\n")

	for _, inputFile := range inputFiles {
		if filepath.Ext(inputFile.Name()) != ".json" {
			continue
		}

		filePath := getFilePath(inputDir, inputFile.Name())

		generatedStructCode, err := generateStructCodeFromFile(filePath)
		if err != nil {
			if shouldStopOnError {
				return errGeneratingStructFromFile(filePath, err)
			}
			fmt.Print(errGeneratingStructFromFile(filePath, err))
			continue
		}

		code.WriteString(generatedStructCode + "\n\n")
	}

	if err := writeToFile(outputFilename, code.String()); err != nil {
		return errWritingOutput(outputFilename, err)
	}

	return nil
}

func generateStructCodeFromFile(jsonFilePath string) (string, error) {
	structName := strings.TrimSuffix(filepath.Base(jsonFilePath), filepath.Ext(jsonFilePath))

	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return "", errOpeningJSONFile(jsonFilePath, err)
	}
	defer jsonFile.Close()

	structFields, err := parseIntoStructFields(jsonFile)
	if err != nil {
		return "", errReadingJSONFile(jsonFilePath, err)
	}

	return generateStructCodeFromFields(structFields, structName), nil
}

func generateStructCodeFromFields(fields []StructField, structName string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for _, field := range fields {
		builder.WriteString(fmt.Sprintf("    %s %s\n", field.Name, field.Type))
	}
	builder.WriteString("}")
	return builder.String()
}

/*  Read & Parse into []StructField the already loaded JSON files */

// IMPROVEMENT: Extend fs.File to have a parseIntoStructFields method
func parseIntoStructFields(jsonFile fs.File) ([]StructField, error) {
	var structFields []StructField
	scanner := bufio.NewScanner(jsonFile)

	for scanner.Scan() {
		structField, err := parseJSONLineIntoStructField(scanner.Text())
		if err != nil || scanner.Err() != nil {
			if err != errParsingShortLine {
				fmt.Println(fmt.Errorf("%w%w", err, scanner.Err()))
			}
			continue
		}

		structFields = append(structFields, structField)
	}

	return structFields, nil
}

func parseJSONLineIntoStructField(line string) (StructField, error) {
	line = strings.Trim(line, " ,{}")

	if len(strings.Trim(line, " ")) < 4 {
		return StructField{}, errParsingShortLine
	}

	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return StructField{}, errParsingJSONLine(line)
	}

	return StructField{
		Name: strings.Trim(parts[0], " \""),
		Type: strings.Trim(parts[1], " \""),
	}, nil
}

/* Simple file helpers */

func readInputFiles(dir string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func writeToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func getFilePath(dir, fileName string) string {
	return filepath.Join(dir, fileName)
}

/* Extra stuff */

type StructField struct {
	Name string
	Type string
}

var (
	errReadingInputFiles = func(err error) error {
		return fmt.Errorf("error reading input files: %v", err)
	}
	errGeneratingStructFromFile = func(filePath string, err error) error {
		return fmt.Errorf("error generating struct from file %s: %v", filePath, err)
	}
	errWritingOutput = func(outputFilename string, err error) error {
		return fmt.Errorf("error writing to %s: %v", outputFilename, err)
	}
	errOpeningJSONFile = func(filePath string, err error) error {
		return fmt.Errorf("error opening JSON file %s: %v", filePath, err)
	}
	errReadingJSONFile = func(filePath string, err error) error {
		return fmt.Errorf("error reading JSON file %s: %v", filePath, err)
	}
	errParsingJSONLine = func(line string) error {
		return fmt.Errorf("error parsing JSON line %s: invalid format", line)
	}
	errParsingShortLine = fmt.Errorf("error parsing JSON line: too short")
)
