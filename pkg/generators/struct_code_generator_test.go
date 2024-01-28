package generators

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var (
	testFileName = "../../in/test.json"
)

func TestStructCodeGeneratorSetSource(t *testing.T) {
	generator := NewStructCodeGenerator("")
	generator.SetSource("test")

	assert.Equal(t, "test", generator.fileName)
}

func TestStructCodeGeneratorRead(t *testing.T) {
	generator := NewStructCodeGenerator(testFileName)
	err := generator.Read()

	assert.Equal(t, nil, err)
	assert.Equal(t, "test", generator.data.Name)
	assert.Equal(t, "name", generator.data.Fields[0].Name)
	assert.Equal(t, "string", generator.data.Fields[0].Type)
	assert.Equal(t, "age", generator.data.Fields[1].Name)
	assert.Equal(t, "int", generator.data.Fields[1].Type)
}

func TestStructCodeGeneratorGenerate(t *testing.T) {
	generator := NewStructCodeGenerator(testFileName)
	generator.Read()
	err := generator.Generate()

	assert.Equal(t, nil, err)
	assert.Equal(t, "type test struct {\n    name string\n    age int\n}", generator.output)
}

func TestStructCodeGeneratorClear(t *testing.T) {
	generator := NewStructCodeGenerator(testFileName)
	generator.Read()
	generator.Generate()
	generator.Clear()

	assert.Equal(t, "", generator.fileName)
	assert.Equal(t, "", generator.output)
	assert.Equal(t, "", generator.data.Name)
	assert.Equal(t, 0, len(generator.data.Fields))
}

func TestStructCodeGeneratorGetOutput(t *testing.T) {
	generator := NewStructCodeGenerator(testFileName)
	generator.Read()
	generator.Generate()

	assert.Equal(t, "type test struct {\n    name string\n    age int\n}", generator.GetOutput())
}
