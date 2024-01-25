package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// StructField examples: { "Name": "ID", "Type": "int" } - { "Name": "Username", "Type": "string" }
type StructField struct {
	Name string
	Type string
}

// processJSONLine input example: "ID": "int",
func processJSONLine(line string) (string, string) {
	line = strings.Trim(line, " ,{}")

	if len(line) < 3 {
		return "", "" // not a valid line
	}

	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return "", "" // not a valid line
	}

	return strings.Trim(parts[0], " \""), strings.Trim(parts[1], " \"")
}

// FromJSONFileToStructFields reads a JSON file and returns an ordered list of StructFields
func FromJSONFileToStructFields(file *os.File) ([]StructField, error) {
	var fields []StructField
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fieldName, fieldType := processJSONLine(scanner.Text())
		fields = append(fields, StructField{Name: fieldName, Type: fieldType})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return fields, nil
}

// GenerateStruct generates Go struct code from a slice of StructField
func GenerateStruct(fields []StructField, structName string) string {
	var structFields []string

	for _, field := range fields {
		structFields = append(structFields, fmt.Sprintf("    %s %s", field.Name, field.Type))
	}

	return fmt.Sprintf("type %s struct {\n%s\n}", structName, strings.Join(structFields, "\n"))
}

// WriteToFile writes the generated struct to a file
func WriteToFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

func ReadInputFiles() []os.DirEntry {
	if files, err := os.ReadDir("in"); err == nil {
		return files
	}

	os.Exit(1)
	return nil
}

const (
	outputFilename = "out/out.go"
)

func WriteOutputFile(generatedCode strings.Builder) {
	if err := WriteToFile(outputFilename, generatedCode.String()); err != nil {
		fmt.Println("Error writing to ", outputFilename, ":", err)
		os.Exit(1)
	}
}

func GetFilePath(inputFile fs.DirEntry) string {
	return filepath.Join("in", inputFile.Name())
}

func main() {

	inputFiles := ReadInputFiles() // from the /in folder, in .json format

	var generatedCode strings.Builder
	generatedCode.WriteString("package main\n\n")

	for _, inputFile := range inputFiles {

		if filepath.Ext(inputFile.Name()) != ".json" {
			continue
		}

		jsonFile, err := os.Open(GetFilePath(inputFile)) // "in/TestModel.json"
		if err != nil {
			fmt.Printf("Error opening JSON file %s: %s\n", GetFilePath(inputFile), err)
			continue // next file
		}
		defer jsonFile.Close()

		structFields, err := FromJSONFileToStructFields(jsonFile)
		if err != nil {
			fmt.Printf("Error reading JSON file %s: %s\n", GetFilePath(inputFile), err)
			continue // next file
		}

		modelName := strings.TrimSuffix(inputFile.Name(), filepath.Ext(inputFile.Name())) // "TestModel"
		structGeneratedCode := GenerateStruct(structFields, modelName)

		generatedCode.WriteString(structGeneratedCode + "\n\n")
	}

	WriteOutputFile(generatedCode)

	fmt.Println("Code successfully written to ", outputFilename)
}
