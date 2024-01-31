package core

import "os"

/* FileWriter Interface */

type FileWriter interface {
	Write(content, fileName string) error
}

/* Struct File Writer */

type StructFileWriter struct{}

func NewStructFileWriter() *StructFileWriter {
	return &StructFileWriter{}
}

// Write writes the string content to a file with the given filename
func (fw *StructFileWriter) Write(content, fileName string) error {
	if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
		return ErrWritingOutput(fileName, err)
	}
	return nil
}
