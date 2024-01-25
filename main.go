package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type StructField struct {
	Name string
	Type string
}

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

func FromJSONFileToStructFields(file fs.File) ([]StructField, error) {
	scanner := bufio.NewScanner(file)
	var fields []StructField
	for scanner.Scan() {
		fieldName, fieldType := processJSONLine(scanner.Text())
		fields = append(fields, StructField{Name: fieldName, Type: fieldType})
	}
	return fields, scanner.Err()
}

func GenerateStruct(fields []StructField, structName string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for _, field := range fields {
		builder.WriteString(fmt.Sprintf("    %s %s\n", field.Name, field.Type))
	}
	builder.WriteString("}")
	return builder.String()
}

func WriteToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func ReadInputFiles(dir string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func GetFilePath(dir, fileName string) string {
	return filepath.Join(dir, fileName)
}

func GenerateStructCodeFromFile(filePath string) (string, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening JSON file %s: %w", filePath, err)
	}
	defer jsonFile.Close()

	structFields, err := FromJSONFileToStructFields(jsonFile)
	if err != nil {
		return "", fmt.Errorf("error reading JSON file %s: %w", filePath, err)
	}

	structName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	return GenerateStruct(structFields, structName), nil
}

const inputDir = "in"
const outputFilename = "out/out.go"

func main() {
	inputFiles, err := ReadInputFiles(inputDir)
	if err != nil {
		fmt.Printf("Error reading input files: %v\n", err)
		return
	}

	var generatedCode strings.Builder
	generatedCode.WriteString("package main\n\n")

	for _, inputFile := range inputFiles {
		if filepath.Ext(inputFile.Name()) != ".json" {
			continue
		}

		filePath := GetFilePath(inputDir, inputFile.Name())
		generatedStructCode, err := GenerateStructCodeFromFile(filePath)
		if err != nil {
			fmt.Printf("Error generating struct from file %s: %v\n", filePath, err)
			continue
		}

		generatedCode.WriteString(generatedStructCode + "\n\n")
	}

	if err := WriteToFile(outputFilename, generatedCode.String()); err != nil {
		fmt.Printf("Error writing to %s: %v\n", outputFilename, err)
		return
	}

	fmt.Printf("Code successfully written to %s\n", outputFilename)
}
