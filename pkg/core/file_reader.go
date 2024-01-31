package core

import (
	"bufio"
	"os"
	"strings"
)

/* FileReader Interface */

type FileReader interface {
	ParseFiles() error
	GetParsedData() []CodeGenerator
}

/* Struct File Reader */

type StructFileReader struct {
	inputFolderPath  string
	parsedStructData []StructData
}

func NewStructFileReader(inputFolderPath string) *StructFileReader {
	return &StructFileReader{
		inputFolderPath:  inputFolderPath,
		parsedStructData: []StructData{},
	}
}

// ParseFiles reads all JSON files in the input folder and parses them into a StructData
func (f *StructFileReader) ParseFiles() error {
	inputFilenames, err := f.getInputFolderFilenames()
	if err != nil {
		return err
	}

	for _, fileName := range inputFilenames {
		jsonFile, err := os.Open(f.inputFolderPath + "/" + fileName)
		if err != nil {
			return ErrOpeningJSONFile(fileName, err)
		}

		f.parsedStructData = append(f.parsedStructData, StructData{Name: fileName, Fields: f.parseIntoStructFields(jsonFile)})

		jsonFile.Close()
	}

	return nil
}

// parseIntoStructFields parses a JSON file into a slice of StructFields
func (f *StructFileReader) parseIntoStructFields(jsonFile *os.File) []StructField {
	var (
		scanner = bufio.NewScanner(jsonFile)
		out     = []StructField{}
	)

	for scanner.Scan() {
		if structField, err := f.parseIntoStructField(scanner.Text()); err == nil {
			out = append(out, structField)
		}
	}

	return out
}

// parseIntoStructField parses a line of JSON into a StructField. Just Name and Type for now
func (f *StructFileReader) parseIntoStructField(line string) (StructField, error) {
	// Remove commas, spaces, and brackets
	line = strings.Trim(line, " ,{}")

	// If line is too short, there is no useful data to parse
	if len(line) < 4 {
		return StructField{}, ErrParsingShortLine
	}

	// Split line into name and type
	// "UserID":"int" -> ["UserID", "int"]
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return StructField{}, ErrParsingJSONLine(line)
	}

	// Remove quotes and backslashes
	return StructField{
		Name: strings.Trim(parts[0], " \""),
		Type: strings.Trim(parts[1], " \""),
	}, nil
}

// getInputFolderFilenames returns a slice with the filenames of the input folder
func (f *StructFileReader) getInputFolderFilenames() ([]string, error) {
	inputFiles, err := os.ReadDir(f.inputFolderPath)
	if err != nil {
		return nil, ErrReadingInputFiles(err)
	}

	var out []string
	for _, inputFile := range inputFiles {
		if strings.HasSuffix(inputFile.Name(), ".json") {
			out = append(out, inputFile.Name())
		}
	}

	return out, nil
}

// GetParsedData returns the parsed data
func (f *StructFileReader) GetParsedData() []CodeGenerator {
	var out []CodeGenerator
	for _, structData := range f.parsedStructData {
		out = append(out, structData)
	}
	return out
}
