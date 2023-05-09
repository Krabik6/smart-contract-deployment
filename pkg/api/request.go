package api

import "encoding/json"

type DeployRequest struct {
	SourceCode           string          `json:"source_code"`
	ConstructorArguments json.RawMessage `json:"constructor_arguments"`
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

type VerifyRequest struct {
	SourceCode           string          `json:"source_code"`
	ContractAddress      string          `json:"contract_address"`
	ContractName         string          `json:"contract_name"`
	ConstructorArguments json.RawMessage `json:"constructor_arguments"`
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
