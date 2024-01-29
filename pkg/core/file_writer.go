package core

import "os"

type FileWriterI interface {
	Write(content, fileName string) error
}

type FileWriter struct {
}

func NewFileWriter() *FileWriter {
	return &FileWriter{}
}

func (fw *FileWriter) Write(content, fileName string) error {
	if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
		return ErrWritingOutput(fileName, err)
	}
	return nil
}
