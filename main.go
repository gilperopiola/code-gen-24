package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// StructField represents a single field in your struct
type StructField struct {
	Name string
	Type string
}

// ReadJSONAndPreserveOrder reads the JSON file and returns an ordered list of StructField
func ReadJSONAndPreserveOrder(filename string) ([]StructField, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fields []StructField
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ,{}")
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue // not a valid line
		}

		name := strings.Trim(parts[0], " \"")
		fieldType := strings.Trim(parts[1], " \"")
		fields = append(fields, StructField{Name: name, Type: fieldType})
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

func main() {
	files, err := ioutil.ReadDir("in")
	if err != nil {
		fmt.Println("Error reading 'in' directory:", err)
		os.Exit(1)
	}

	var fullCode strings.Builder
	fullCode.WriteString("package main\n\n")

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			jsonFile := filepath.Join("in", file.Name())
			structName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

			fields, err := ReadJSONAndPreserveOrder(jsonFile)
			if err != nil {
				fmt.Printf("Error reading JSON file %s: %s\n", jsonFile, err)
				continue // Skip to the next file
			}

			structCode := GenerateStruct(fields, structName)
			fullCode.WriteString(structCode + "\n\n")
		}
	}

	outputFile := "out/out.go"
	if err := WriteToFile(outputFile, fullCode.String()); err != nil {
		fmt.Println("Error writing to", outputFile, ":", err)
		os.Exit(1)
	}

	fmt.Println("Structs successfully written to", outputFile)
}
