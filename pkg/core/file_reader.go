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

func (f *StructFileReader) GetParsedData() []CodeGenerator {
	var out []CodeGenerator
	for _, structData := range f.parsedStructData {
		out = append(out, structData)
	}
	return out
}

func (f *StructFileReader) ParseFiles() error {
	inputFilenames, err := f.getInputFolderFilenames()
	if err != nil {
		return err
	}

	for _, fileName := range inputFilenames {
		jsonFile, err := os.Open(fileName)
		if err != nil {
			return ErrOpeningJSONFile(fileName, err)
		}

		f.parsedStructData = append(f.parsedStructData, f.parseIntoStructData(jsonFile))

		jsonFile.Close()
	}

	return nil
}

func (f *StructFileReader) parseIntoStructData(jsonFile *os.File) StructData {
	return StructData{
		Name:   strings.TrimPrefix(strings.TrimSuffix(jsonFile.Name(), ".json"), f.inputFolderPath+"/"),
		Fields: f.parseIntoStructFields(jsonFile),
	}
}

func (f *StructFileReader) parseIntoStructFields(jsonFile *os.File) []StructField {
	var (
		out     = []StructField{}
		scanner = bufio.NewScanner(jsonFile)
	)

	for scanner.Scan() {
		if structField, err := f.parseIntoStructField(scanner.Text()); err == nil {
			out = append(out, structField)
		}
		continue
	}

	return out
}

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

func (f *StructFileReader) getInputFolderFilenames() ([]string, error) {
	var out []string

	inputFiles, err := os.ReadDir(f.inputFolderPath)
	if err != nil {
		return nil, ErrReadingInputFiles(err)
	}

	for _, inputFile := range inputFiles {
		if strings.HasSuffix(inputFile.Name(), ".json") {
			out = append(out, f.inputFolderPath+"/"+inputFile.Name())
		}
	}

	return out, nil
}
