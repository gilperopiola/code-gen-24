package pkg

// CodeGenerator is the interface to be implemented by each struct holding data used to generate code
type CodeGenerator interface {
	ParseData(data []byte) error
	GenerateCode() string
}
