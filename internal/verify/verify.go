package verify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	EncodeConstructorArgs(sourceCode string, args ...interface{}) ([]byte, error)
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
	//constructorArguments type and type of each element
	log.Println("constructorArguments type:", fmt.Sprintf("%T", constructorArguments))
	for i, v := range constructorArguments {
		log.Println("constructorArguments element type:", fmt.Sprintf("%T", v), "element:", i)
	}

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
		args, err := v.Compiler.EncodeConstructorArgs(sourceCode, []interface{}{uint(8712354)}...)
		if err != nil {
			return err
		}
		log.Println("args:", args)
		log.Println("encode args:", hexutil.Encode(args))
		//[]interface{} to string

		params["constructorarguments"] = hexutil.Encode(args)
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
