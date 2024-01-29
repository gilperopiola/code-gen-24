package core

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileReaderI interface {
	SetInputFolder(inputFolder string)
	ReadInputFolder() error
	GetParsedStructData() []Generable
}

type FileReader struct {
	inputFolderPath string
	inputFilenames  []string

	parsedStructData map[string]StructData
}

func NewFileReader(inputFolderPath string) *FileReader {
	return &FileReader{
		inputFolderPath:  inputFolderPath,
		inputFilenames:   []string{},
		parsedStructData: make(map[string]StructData),
	}
}

func (f *FileReader) SetInputFolder(inputFolderPath string) {
	f.inputFolderPath = inputFolderPath
}

func (f *FileReader) GetParsedStructData() []Generable {
	var parsedStructData []Generable

	for _, structData := range f.parsedStructData {
		parsedStructData = append(parsedStructData, structData)
	}

	return parsedStructData
}

func (f *FileReader) ReadInputFolder() error {
	f.loadFilenames()

	for _, fileName := range f.inputFilenames {
		fileName = f.inputFolderPath + "/" + fileName

		f.parsedStructData[fileName] = StructData{
			Name:   strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName)),
			Fields: []StructField{},
		}

		jsonFile, err := os.Open(fileName)
		if err != nil {
			return ErrOpeningJSONFile(fileName, err)
		}

		f.parseIntoStructFields(jsonFile, fileName)

		jsonFile.Close()
	}

	return nil
}

func (f *FileReader) parseIntoStructFields(jsonFile *os.File, fileName string) {
	scanner := bufio.NewScanner(jsonFile)

	for scanner.Scan() {

		structField, err := f.parseJSONLineIntoStructField(scanner.Text())
		if err != nil {
			if err != ErrParsingShortLine {
				fmt.Println(err)
			}
			continue
		}

		fields := append(f.parsedStructData[fileName].Fields, structField)
		f.parsedStructData[fileName] = StructData{
			Name:   f.parsedStructData[fileName].Name,
			Fields: fields,
		}

	}

	if scanner.Err() != nil {
		fmt.Println(fmt.Errorf("error scanning JSON file %s: %v", fileName, scanner.Err()))
	}
}

func (f *FileReader) parseJSONLineIntoStructField(line string) (StructField, error) {
	line = strings.Trim(line, " ,{}")

	if len(strings.Trim(line, " ")) < 4 {
		return StructField{}, ErrParsingShortLine
	}

	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return StructField{}, ErrParsingJSONLine(line)
	}

	return StructField{
		Name: strings.Trim(parts[0], " \""),
		Type: strings.Trim(parts[1], " \""),
	}, nil
}

func (f *FileReader) loadFilenames() {
	inputFolderFiles, err := os.ReadDir(f.inputFolderPath)
	if err != nil {
		fmt.Println(ErrReadingInputFiles(err))
		return
	}

	var fileNames []string

	for _, inputFile := range inputFolderFiles {
		fileName := inputFile.Name()

		if !strings.HasSuffix(fileName, ".json") {
			continue
		}

		fileNames = append(fileNames, fileName)
	}

	f.inputFilenames = fileNames
}
