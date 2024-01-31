package core

import "fmt"

var (
	ErrReadingInputFiles = func(err error) error {
		return fmt.Errorf("error reading input files: %v", err)
	}
	ErrWritingOutput = func(outputFilename string, err error) error {
		return fmt.Errorf("error writing to %s: %v", outputFilename, err)
	}
	ErrOpeningJSONFile = func(filePath string, err error) error {
		return fmt.Errorf("error opening JSON file %s: %v", filePath, err)
	}
	ErrParsingJSONLine = func(line string) error {
		return fmt.Errorf("error parsing JSON line %s: invalid format", line)
	}
	ErrParsingShortLine = fmt.Errorf("error parsing JSON line: too short")
)
