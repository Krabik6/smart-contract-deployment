package verify

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
	"strconv"
)

type Verifier struct {
	ArgsEncoder ArgsEncoder
}

func NewVerifier(argsEncoder ArgsEncoder) *Verifier {
	return &Verifier{
		ArgsEncoder: argsEncoder,
	}
}

//
//// Compiler is the interface that wraps the Compile method.
//type Compiler interface {
//	EncodeConstructorArgs(sourceCode string, args ...interface{}) ([]byte, error)
//}
//
//type CompilerJson interface {
//	GetBytecode(inputJSON []byte, contractPath, contractName string) ([]byte, error)
//	GetAbi(inputJSON []byte, contractPath, contractName string) (abi.ABI, error)
//}

type ArgsEncoder interface {
	EncodeConstructorArgs(abi abi.ABI, args ...interface{}) ([]byte, error)
}

func BoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

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

func (v *Verifier) Verify(abi abi.ABI, params Params, constructorArguments ...interface{}) error {
	optimizeStr := "0"
	if params.OptimizationUsed != nil {
		optimizeStr = BoolToString(*params.OptimizationUsed)
	}

	log.Println("sourceCode:", params.SourceCode)
	log.Println("format:", params.CodeFormat)

	_params := map[string]string{
		"apikey":           "AFEMDPHAWXPHKI8SQJK9AS77UIAZN9NGCN",
		"module":           "contract",
		"action":           "verifysourcecode",
		"contractaddress":  params.ContractAddress,
		"sourcecode":       params.SourceCode,
		"codeformat":       params.CodeFormat,
		"contractname":     params.ContractName,
		"compilerversion":  params.CompilerVersion,
		"optimizationUsed": optimizeStr,
		"licenseType":      strconv.Itoa(*params.LicenseType),
	}

	// if optimize == true
	if params.OptimizationUsed != nil {
		if *params.OptimizationUsed {
			runsStr := strconv.Itoa(*params.Runs)
			_params["runs"] = runsStr
		}
	}

	// if constructorArguments len > 0
	if len(constructorArguments) != 0 {
		args, err := v.ArgsEncoder.EncodeConstructorArgs(abi, constructorArguments...)
		if err != nil {
			return err
		}
		hexArgsWithout0x := hex.EncodeToString(args)
		_params["constructorArguements"] = hexArgsWithout0x
	}

	formData := url.Values{}
	for k, v := range _params {
		formData.Set(k, v)
	}

	reqBody := bytes.NewBufferString(formData.Encode())

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api-testnet.polygonscan.com/api", reqBody)
	if err != nil {
		debug.PrintStack()
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// print request body for debug
	log.Println("request body:", reqBody)
	// print request body constructorArguments for debug
	log.Println("request body constructorArguments:", _params["constructorArguements"])

	res, err := client.Do(req)
	if err != nil {
		debug.PrintStack()
		return err
	}

	defer res.Body.Close()
	var result map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		debug.PrintStack()
		return err
	}

	if status, ok := result["status"].(string); !ok || status != "1" {
		return errors.New(result["result"].(string))
	}
	log.Println("result:", result["result"])
	log.Println("message:", result["message"])
	log.Println("status:", result["status"])

	return nil
}

//{
//"status": "1",
//"message": "OK",
//"result": "d9vrjsemlffmhmwmxpuhs1adbvtth2dszui52rergfjzwgxzvt"
//}
