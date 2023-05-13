package api

import "encoding/json"

type DeployRequest struct {
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

//Verify(contractAddress, sourceCode, contractName, constructorArguments, licenseType, compilerversion string, optimize bool, runs int) error

type VerifyRequest struct {
	ConstructorArguments json.RawMessage `json:"arguments"`
	APIKey               string          `json:"apikey"`
	ContractAddress      string          `json:"contractaddress"`
	SourceCode           string          `json:"sourceCode"`
	CodeFormat           string          `json:"codeformat"`
	ContractName         string          `json:"contractname"`
	CompilerVersion      string          `json:"compilerversion"`
	OptimizationUsed     *bool           `json:"optimizationUsed,omitempty"`
	Runs                 *int            `json:"runs,omitempty"`
	EVMVersion           *string         `json:"evmversion,omitempty"`
	LicenseType          *int            `json:"licenseType,omitempty"`
	LibraryName1         *string         `json:"libraryname1,omitempty"`
	LibraryAddress1      *string         `json:"libraryaddress1,omitempty"`
	LibraryName2         *string         `json:"libraryname2,omitempty"`
	LibraryAddress2      *string         `json:"libraryaddress2,omitempty"`
	LibraryName3         *string         `json:"libraryname3,omitempty"`
	LibraryAddress3      *string         `json:"libraryaddress3,omitempty"`
	LibraryName4         *string         `json:"libraryname4,omitempty"`
	LibraryAddress4      *string         `json:"libraryaddress4,omitempty"`
	LibraryName5         *string         `json:"libraryname5,omitempty"`
	LibraryAddress5      *string         `json:"libraryaddress5,omitempty"`
	LibraryName6         *string         `json:"libraryname6,omitempty"`
	LibraryAddress6      *string         `json:"libraryaddress6,omitempty"`
	LibraryName7         *string         `json:"libraryname7,omitempty"`
	LibraryAddress7      *string         `json:"libraryaddress7,omitempty"`
	LibraryName8         *string         `json:"libraryname8,omitempty"`
	LibraryAddress8      *string         `json:"libraryaddress8,omitempty"`
	LibraryName9         *string         `json:"libraryname9,omitempty"`
	LibraryAddress9      *string         `json:"libraryaddress9,omitempty"`
	LibraryName10        *string         `json:"libraryname10,omitempty"`
	LibraryAddress10     *string         `json:"libraryaddress10,omitempty"`
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
