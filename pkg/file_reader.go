package pkg

import (
	"os"
	"path/filepath"
	"strings"
)

/* FileReader Interface */

type FileReader interface {
	ParseFiles() error
	GetParsedData() []CodeGenerator
}

/* Struct File Reader */

type StructFileReader struct {
	inputFolder      string
	parsedStructData []StructData
}

func NewStructFileReader(inputFolder string) *StructFileReader {
	return &StructFileReader{
		inputFolder:      inputFolder,
		parsedStructData: []StructData{},
	}
}

// ParseFiles reads all .json files in the input folder and parses them into a []StructData
func (f *StructFileReader) ParseFiles() error {

	// Get filenames
	inputFilenames, err := f.getInputJSONFilenames()
	if err != nil {
		return err
	}

	for _, fileName := range inputFilenames {

		// Read file
		fileData, err := os.ReadFile(filepath.Join(f.inputFolder, fileName))
		if err != nil {
			return ErrReadingFile(fileName, err)
		}

		// Create StructData with Name
		structData := StructData{Name: strings.TrimSuffix(fileName, ".json")}

		// Parse JSON contents into Fields
		if err = structData.ParseData(fileData); err != nil {
			return err
		}

		f.parsedStructData = append(f.parsedStructData, structData)
	}

	return nil
}

// GetParsedData returns the parsed data
func (f *StructFileReader) GetParsedData() []CodeGenerator {
	var out []CodeGenerator
	for _, structData := range f.parsedStructData {
		structData := structData // This changes the pointer address, needed for the line below to work correctly
		out = append(out, &structData)
	}
	return out
}

// getInputJSONFilenames returns a slice with the names of the .json files in the input folder
func (f *StructFileReader) getInputJSONFilenames() ([]string, error) {

	// Get input folder files
	entries, err := os.ReadDir(f.inputFolder)
	if err != nil {
		return nil, ErrReadingInputFolder(err)
	}

	// Leave only the .json ones
	var out []string
	for _, file := range entries {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			out = append(out, file.Name())
		}
	}

	return out, nil
}
