package handler

import "github.com/ethereum/go-ethereum/accounts/abi"

// Deployer is the interface that wraps the Deploy method.
type Deployer interface {
	Deploy(sourceCode string, optimize bool, runs int, args ...interface{}) (string, error)
	EstimateGas(sourceCode string, optimize bool, runs int, args ...interface{}) (int, error)
}

// Compiler is the interface that wraps the Compile method.
type Compiler interface {
	GetAbi(sourceCode string) (abi.ABI, error)
	GetBytecode(sourceCode string, optimize bool, runs int) ([]byte, error)
	//ConvertAndCheckArgs(args []interface{}, contractAbiJson *abi.ABI) ([]interface{}, error)
	EncodeConstructorArgs(sourceCode string, args ...interface{}) ([]byte, error)
	EncodeFunctionCall(sourceCode string, functionName string, args ...interface{}) ([]byte, error)
	GetAbiFromString(abiString string) (abi.ABI, error)
}

// Verifier is the interface that wraps the Verify method.
type Verifier interface {
	Verify(contractAddress, sourceCode, contractName, licenseType, compilerversion string, optimize bool, runs int, constructorArguments ...interface{}) error
}
