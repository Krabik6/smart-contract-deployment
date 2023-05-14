package verify

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
	"reflect"
	"strconv"
)

func BoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
func (v *Verifier) logParams(params Params) {
	optimizeStr := "0"
	if params.OptimizationUsed != nil {
		optimizeStr = BoolToString(*params.OptimizationUsed)
	}

	log.Println("source code: ", params.SourceCode)
	log.Println("contract name: ", params.ContractName)
	log.Println("compiler version: ", params.CompilerVersion)
	log.Println("optimization used: ", optimizeStr)
	log.Println("license type: ", params.LicenseType)
	log.Println("code format: ", params.CodeFormat)
}

func (v *Verifier) prepareParams(abi abi.ABI, params Params, constructorArguments ...interface{}) (map[string]string, error) {
	optimizeStr := "0"
	if params.OptimizationUsed != nil {
		optimizeStr = BoolToString(*params.OptimizationUsed)
	}

	_params := map[string]string{
		"apikey":           params.APIKey,
		"module":           "contract",
		"action":           "verifysourcecode",
		"contractaddress":  params.ContractAddress,
		"sourcecode":       params.SourceCode,
		"codeformat":       params.CodeFormat,
		"contractname":     params.ContractName,
		"compilerversion":  params.CompilerVersion,
		"optimizationUsed": optimizeStr,
		"licenseType":      strconv.Itoa(params.LicenseType),
	}

	if params.OptimizationUsed != nil && *params.OptimizationUsed && params.Runs != nil {
		_params["runs"] = strconv.Itoa(*params.Runs)
	}

	if params.EVMVersion != nil {
		_params["evmversion"] = *params.EVMVersion
	}

	for i := 1; i <= 10; i++ {
		libraryNameField := fmt.Sprintf("LibraryName%d", i)
		libraryAddressField := fmt.Sprintf("LibraryAddress%d", i)

		libraryName := reflect.ValueOf(&params).Elem().FieldByName(libraryNameField).Interface().(*string)
		libraryAddress := reflect.ValueOf(&params).Elem().FieldByName(libraryAddressField).Interface().(*string)

		if libraryName != nil && libraryAddress != nil {
			_params["libraryname"+strconv.Itoa(i)] = *libraryName
			_params["libraryaddress"+strconv.Itoa(i)] = *libraryAddress
		}
	}

	if len(constructorArguments) != 0 {
		args, err := v.ArgsEncoder.EncodeConstructorArgs(abi, constructorArguments...)
		if err != nil {
			return nil, err
		}
		_params["constructorArguements"] = hex.EncodeToString(args)
	}

	return _params, nil
}
