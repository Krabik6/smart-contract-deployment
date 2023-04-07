package handler

// Deployer is the interface that wraps the Deploy method.
type Deployer interface {
	Deploy(sourceCode string, constructorArguments []string) (string, error)
}

// Compiler is the interface that wraps the Compile method.
type Compiler interface {
	GetABI(sourceCode string) (string, error)
	GetBytecode(sourceCode string) (string, error)
	GetABIAndBytecode(sourceCode string) (string, string, error)
}
