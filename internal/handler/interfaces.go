package handler

import (
	"github.com/Krabik6/smart-contract-deployment/internal/verify"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Deployer is the interface that wraps the Deploy method.
type Deployer interface {
	Deploy(bytecode []byte, abi abi.ABI, args ...interface{}) (string, error)
	EstimateGas(sourceCode string, optimize bool, runs int, args ...interface{}) (int, error)
}

// Compiler is the interface that wraps the Compile method.
type Compiler interface {
	GetAbi(sourceCode string) (abi.ABI, error)
	GetBytecode(sourceCode string, optimize bool, runs int) ([]byte, error)
	GetAbiFromString(abiString string) (abi.ABI, error)
}

type CompilerJson interface {
	GetBytecode(inputJSON []byte, contractPath, contractName string) ([]byte, error)
	GetAbi(inputJSON []byte, contractPath, contractName string) (abi.ABI, error)
}

type ArgsEncoder interface {
	EncodeConstructorArgs(abi abi.ABI, args ...interface{}) ([]byte, error)
	EncodeFunctionCall(abi abi.ABI, functionName string, args ...interface{}) ([]byte, error)
}

// Verifier is the interface that wraps the Verify method.
type Verifier interface {
	Verify(networkName string, network verify.Network, abi abi.ABI, params verify.Params, constructorArguments ...interface{}) error
}

type InputGenerator interface {
	GenerateJSONInput(mainSolPath string, optimize bool, optimizeRuns int) ([]byte, string, error)
}
