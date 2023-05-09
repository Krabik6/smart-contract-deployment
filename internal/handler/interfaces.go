package handler

import "github.com/ethereum/go-ethereum/accounts/abi"

// Deployer is the interface that wraps the Deploy method.
type Deployer interface {
	Deploy(sourceCode string, optimize bool, runs int, args ...interface{}) (string, error)
	EstimateGas(sourceCode string, args ...interface{}) (int, error)
}

// Compiler is the interface that wraps the Compile method.
type Compiler interface {
	GetAbi(sourceCode string) (abi.ABI, error)
	GetBytecode(sourceCode string, optimize bool, runs int) ([]byte, error)
}
