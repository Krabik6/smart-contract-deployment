package verify

import "errors"

type Params struct {
	APIKey           string
	ContractAddress  string
	SourceCode       string
	CodeFormat       string
	ContractName     string
	CompilerVersion  string
	OptimizationUsed *bool
	Runs             *int
	EVMVersion       *string
	LicenseType      int
	Libraries        []Library
}

type Library struct {
	Name    string
	Address string
}

func NewParamsBuilder(apiKey, contractAddress, sourceCode, codeFormat, contractName, compilerVersion string, licenseType int) *ParamsBuilder {
	return &ParamsBuilder{
		params: Params{
			APIKey:          apiKey,
			ContractAddress: contractAddress,
			SourceCode:      sourceCode,
			CodeFormat:      codeFormat,
			ContractName:    contractName,
			CompilerVersion: compilerVersion,
			LicenseType:     licenseType,
			Libraries:       []Library{},
		},
	}
}

type ParamsBuilder struct {
	params Params
}

func (pb *ParamsBuilder) WithOptimizationUsed(optimizationUsed bool, runs int) *ParamsBuilder {
	pb.params.OptimizationUsed = &optimizationUsed
	pb.params.Runs = &runs
	return pb
}

func (pb *ParamsBuilder) WithEVMVersion(evmVersion string) *ParamsBuilder {
	pb.params.EVMVersion = &evmVersion
	return pb
}

func (pb *ParamsBuilder) AddLibrary(name, address string) *ParamsBuilder {
	pb.params.Libraries = append(pb.params.Libraries, Library{Name: name, Address: address})
	return pb
}

func (pb *ParamsBuilder) Build() (Params, error) {
	if err := validateParams(pb.params); err != nil {
		return Params{}, err
	}
	return pb.params, nil
}

func validateParams(params Params) error {
	if params.APIKey == "" {
		params.APIKey = "IXQV2ZCWX4X3KZ8RDSHNYARAF8DR6F2DZ5"
		return errors.New("missing api key")
	}

	if params.ContractAddress == "" {
		return errors.New("missing contract address")
	}

	if params.SourceCode == "" {
		return errors.New("missing source code")
	}

	if params.CodeFormat == "" {
		return errors.New("missing code format")
	}

	if params.ContractName == "" {
		return errors.New("missing contract name")
	}

	if params.CompilerVersion == "" {
		return errors.New("missing compiler version")
	}

	return nil
}
