package core

import (
	"bufio"
	"os"
	"strings"
)

type FileReaderI interface {
	ParseFiles() error
	GetParsedData() []Generable
}

type StructFileReader struct {
	inputFolderPath string
	inputFilenames  []string

	parsedStructData []StructData
}

func NewStructFileReader(inputFolderPath string) *StructFileReader {
	return &StructFileReader{
		inputFolderPath: inputFolderPath,
		inputFilenames:  []string{},

		parsedStructData: []StructData{},
	}
}

func (f *StructFileReader) GetParsedData() []Generable {
	var out []Generable
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

		f.parsedStructData = append(f.parsedStructData, f.parseJSONFileIntoStructData(jsonFile, fileName))

		jsonFile.Close()
	}

	return nil
}

func (f *StructFileReader) parseJSONFileIntoStructData(jsonFile *os.File, fileName string) StructData {
	return StructData{
		Name:   strings.TrimSuffix(fileName, ".json"),
		Fields: f.parseJSONFileIntoStructFields(jsonFile),
	}
}

func (f *StructFileReader) parseJSONFileIntoStructFields(jsonFile *os.File) []StructField {
	var (
		out     = []StructField{}
		scanner = bufio.NewScanner(jsonFile)
	)

	for scanner.Scan() {
		if structField, err := f.parseJSONLineIntoStructField(scanner.Text()); err == nil {
			out = append(out, structField)
		}
		continue
	}

	return out
}

func (f *StructFileReader) parseJSONLineIntoStructField(line string) (StructField, error) {

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
	inputFiles, err := os.ReadDir(f.inputFolderPath)
	if err != nil {
		return nil, ErrReadingInputFiles(err)
	}

	for _, inputFile := range inputFiles {
		if strings.HasSuffix(inputFile.Name(), ".json") {
			f.inputFilenames = append(f.inputFilenames, f.inputFolderPath+"/"+inputFile.Name())
		}
	}

	return f.inputFilenames, nil
}
