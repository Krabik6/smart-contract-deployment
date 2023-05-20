package api

import "encoding/json"

type DeployRequest struct {
	PrivateKey  string `json:"privatekey"`
	Provider    string `json:"provider"`
	NetworkName string `json:"network"`

	SourceCode           string          `json:"source_code"`
	ConstructorArguments json.RawMessage `json:"arguments"`
	Optimize             bool            `json:"optimize"`
	Runs                 int             `json:"runs"`
	ContractName         string          `json:"contract_name"`
}

type AbiRequest struct {
	SourceCode string `json:"source_code"`
}

type BytecodeRequest struct {
	SourceCode string `json:"source_code"`
	Optimize   bool   `json:"optimize"`
	Runs       int    `json:"runs"`
}

// Verify(contractAddress, sourceCode, contractName, constructorArguments, licenseType, compilerversion string, optimize bool, runs int) error
type Library struct {
	Name    *string `json:"name,omitempty"`
	Address *string `json:"address,omitempty"`
}

type VerifyRequest struct {
	APIKey      string `json:"apikey"`
	Url         string `json:"url"`
	NetworkName string `json:"network"`

	ConstructorArguments json.RawMessage `json:"arguments"`
	ContractAddress      string          `json:"contractaddress"`
	SourceCode           string          `json:"sourceCode"`
	CodeFormat           string          `json:"codeformat"`
	ContractName         string          `json:"contractname"`
	CompilerVersion      string          `json:"compilerversion"`
	OptimizationUsed     *bool           `json:"optimizationUsed,omitempty"`
	Runs                 *int            `json:"runs,omitempty"`
	EVMVersion           *string         `json:"evmversion,omitempty"`
	LicenseType          int             `json:"licenseType"`
	Libraries            []Library       `json:"libraries,omitempty"`
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

//
//type Params struct {
//
//}
