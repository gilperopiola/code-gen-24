package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsd(t *testing.T) {
	a := GenerateStruct([]StructField{{Name: "A", Type: "int"}, {Name: "B", Type: "string"}}, "TestModel")
	assert.NotEqual(t, "", a)
}
