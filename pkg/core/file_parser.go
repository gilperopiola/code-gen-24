package core

/* FileParser Interface */

type FileParser interface {
}

/* JSON File Parser */

type JSONFileParser struct {
}

func NewJSONFileParser() *JSONFileParser {
	return &JSONFileParser{}
}
