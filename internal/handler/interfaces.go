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
	//
	//EncodeConstructorArgs(sourceCode string, args ...interface{}) ([]byte, error)
	//EncodeFunctionCall(sourceCode string, functionName string, args ...interface{}) ([]byte, error)
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
	Verify(abi abi.ABI, params verify.Params, constructorArguments ...interface{}) error
}

type InputGenerator interface {
	GenerateJSONInput(mainSolPath string, optimize bool, optimizeRuns int) ([]byte, string, error)
}

type Params struct {
	//APIKey           string  `json:"apikey"`
	ContractAddress  string  `json:"contractaddress"`
	SourceCode       string  `json:"sourceCode"`
	CodeFormat       string  `json:"codeformat"`
	ContractName     string  `json:"contractname"`
	CompilerVersion  string  `json:"compilerversion"`
	OptimizationUsed *bool   `json:"optimizationUsed,omitempty"`
	Runs             *int    `json:"runs,omitempty"`
	EVMVersion       *string `json:"evmversion,omitempty"`
	LicenseType      *int    `json:"licenseType,omitempty"`
	LibraryName1     *string `json:"libraryname1,omitempty"`
	LibraryAddress1  *string `json:"libraryaddress1,omitempty"`
	LibraryName2     *string `json:"libraryname2,omitempty"`
	LibraryAddress2  *string `json:"libraryaddress2,omitempty"`
	LibraryName3     *string `json:"libraryname3,omitempty"`
	LibraryAddress3  *string `json:"libraryaddress3,omitempty"`
	LibraryName4     *string `json:"libraryname4,omitempty"`
	LibraryAddress4  *string `json:"libraryaddress4,omitempty"`
	LibraryName5     *string `json:"libraryname5,omitempty"`
	LibraryAddress5  *string `json:"libraryaddress5,omitempty"`
	LibraryName6     *string `json:"libraryname6,omitempty"`
	LibraryAddress6  *string `json:"libraryaddress6,omitempty"`
	LibraryName7     *string `json:"libraryname7,omitempty"`
	LibraryAddress7  *string `json:"libraryaddress7,omitempty"`
	LibraryName8     *string `json:"libraryname8,omitempty"`
	LibraryAddress8  *string `json:"libraryaddress8,omitempty"`
	LibraryName9     *string `json:"libraryname9,omitempty"`
	LibraryAddress9  *string `json:"libraryaddress9,omitempty"`
	LibraryName10    *string `json:"libraryname10,omitempty"`
	LibraryAddress10 *string `json:"libraryaddress10,omitempty"`
}
