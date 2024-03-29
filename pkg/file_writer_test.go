package pkg

import (
	"os"
	"testing"
)

func TestFileWriterWrite(t *testing.T) {
	if err := NewStructFileWriter().Write("test", "test.txt"); err != nil {
		t.Errorf("expected no error, got %v", err)
	} else {
		os.Remove("test.txt")
	}
}
