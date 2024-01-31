package core

import "testing"

func TestStructFileReaderParseFolder(t *testing.T) {
	if err := NewStructFileReader("../../in").ParseFiles(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := NewStructFileReader("nonexistent").ParseFiles(); err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestStructFileReaderGetParsedData(t *testing.T) {
	reader := NewStructFileReader("../../in")
	reader.ParseFiles()
	if len(reader.GetParsedData()) == 0 {
		t.Errorf("expected non-empty array, got empty")
	}
}
