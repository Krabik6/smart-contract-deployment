package api

import "encoding/json"

type DeployRequest struct {
	SourceCode           string          `json:"source_code"`
	ConstructorArguments json.RawMessage `json:"arguments"`
	Optimize             bool            `json:"optimize"`
	Runs                 int             `json:"runs"`
}

type AbiRequest struct {
	SourceCode string `json:"source_code"`
}

type BytecodeRequest struct {
	SourceCode string `json:"source_code"`
	Optimize   bool   `json:"optimize"`
	Runs       int    `json:"runs"`
}

//Verify(contractAddress, sourceCode, contractName, constructorArguments, licenseType, compilerversion string, optimize bool, runs int) error

type VerifyRequest struct {
	ContractAddress      string          `json:"contract_address"`
	SourceCode           string          `json:"source_code"`
	ContractName         string          `json:"contract_name"`
	ConstructorArguments json.RawMessage `json:"arguments"`
	LicenseType          string          `json:"license_type"`
	Compilerversion      string          `json:"compilerversion"`
	Optimize             bool            `json:"optimize"`
	Runs                 int             `json:"runs"`
}

type EncodeFunctionCallRequest struct {
	SourceCode   string          `json:"source_code"`
	FunctionName string          `json:"function_name"`
	Arguments    json.RawMessage `json:"arguments"`
}

type EncodeConstructorArgsRequest struct {
	SourceCode string          `json:"source_code"`
	Arguments  json.RawMessage `json:"arguments"`
}
