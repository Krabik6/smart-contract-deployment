package verify

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"runtime/debug"
)

func Verify(contractAddress, sourceCode, contractName, constructorArguments string) error {

	params := map[string]string{
		"apikey":               "AFEMDPHAWXPHKI8SQJK9AS77UIAZN9NGCN",
		"module":               "contract",
		"action":               "verifysourcecode",
		"contractaddress":      contractAddress,
		"sourcecode":           sourceCode,
		"codeformat":           "solidity-single-file",
		"contractname":         contractName,
		"compilerversion":      "v0.8.19+commit.7dd6d404",
		"optimizationUsed":     "1",
		"runs":                 "200",
		"constructorarguments": constructorArguments,
		"evmversion":           "petersburg",
		"licenseType":          "5",
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
		return errors.New("verification failed")
	}

	return nil
}
