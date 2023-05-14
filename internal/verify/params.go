package verify

import "errors"

type Params struct {
	APIKey           string  `json:"apikey"`
	ContractAddress  string  `json:"contractaddress"`
	SourceCode       string  `json:"sourceCode"`
	CodeFormat       string  `json:"codeformat"`
	ContractName     string  `json:"contractname"`
	CompilerVersion  string  `json:"compilerversion"`
	OptimizationUsed *bool   `json:"optimizationUsed,omitempty"`
	Runs             *int    `json:"runs,omitempty"`
	EVMVersion       *string `json:"evmversion,omitempty"`
	LicenseType      int     `json:"licenseType"`
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

func (v *Verifier) validateParams(params Params) error {
	if params.APIKey == "" {
		params.APIKey = "IXQV2ZCWX4X3KZ8RDSHNYARAF8DR6F2DZ5"
		//return errors.New("missing API key")
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
