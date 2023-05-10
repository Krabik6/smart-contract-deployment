package verify

import (
	"bytes"
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
	Compiler Compiler
}

func NewVerifier(compiler Compiler) *Verifier {
	return &Verifier{
		Compiler: compiler,
	}
}

// Compiler is the interface that wraps the Compile method.
type Compiler interface {
	//ConvertAndCheckArgs(args []interface{}, contractAbiJson *abi.ABI) ([]interface{}, error)
	EncodeConstructorArgs(sourceCode string, args ...interface{}) ([]byte, error)
	EncodeFunctionCall(sourceCode string, functionName string, args ...interface{}) ([]byte, error)
	GetAbiFromString(abiString string) (abi.ABI, error)
}

func BoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func (v *Verifier) Verify(contractAddress, sourceCode, contractName, licenseType, compilerversion string, optimize bool, runs int, constructorArguments []interface{}) error {
	optimizeStr := BoolToString(optimize)
	// print all args verify function
	log.Println("contractAddress:", contractAddress)
	log.Println("sourceCode:", sourceCode)
	log.Println("contractName:", contractName)
	log.Println("licenseType:", licenseType)
	log.Println("compilerversion:", compilerversion)
	log.Println("optimize:", optimize)
	log.Println("runs:", runs)
	log.Println("constructorArguments:", constructorArguments)

	params := map[string]string{
		"apikey":           "AFEMDPHAWXPHKI8SQJK9AS77UIAZN9NGCN",
		"module":           "contract",
		"action":           "verifysourcecode",
		"contractaddress":  contractAddress,
		"sourcecode":       sourceCode,
		"codeformat":       "solidity-single-file",
		"contractname":     contractName,
		"compilerversion":  compilerversion,
		"optimizationUsed": optimizeStr,
		"licenseType":      licenseType,
	}

	// if optimize == true
	if optimize {
		runsStr := strconv.Itoa(runs)
		params["runs"] = runsStr
	}

	// if constructorArguments len > 0
	if len(constructorArguments) != 0 {
		args, err := v.Compiler.EncodeConstructorArgs(sourceCode, constructorArguments)
		if err != nil {
			return err
		}
		strArgs := string(args)
		params["constructorarguments"] = strArgs
	}

	formData := url.Values{}
	for k, v := range params {
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

	//print req body
	// print and parse to readable human format
	log.Printf("Request: Method=%s, URL=%s, Headers=%v, Body=%s", req.Method, req.URL.String(), req.Header, reqBody.String())

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

	return nil
}
