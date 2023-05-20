package verify

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
)

func (v *Verifier) prepareRequest(baseURL string, params map[string]string) (*http.Request, error) {
	formData := url.Values{}
	for k, v := range params {
		formData.Set(k, v)
	}

	log.Println("BaseURL:", baseURL)

	reqBody := bytes.NewBufferString(formData.Encode())

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	req, err := http.NewRequest("POST", "https://api-testnet.polygonscan.com/api", reqBody)
	if err != nil {
		debug.PrintStack()
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// print request body for debug
	log.Println("request body:", reqBody)
	// print request body constructorArguments for debug
	log.Println("request body constructorArguments:", params["constructorArguements"])

	return req, nil
}

func (v *Verifier) sendRequestAndParseResponse(req *http.Request) error {
	res, err := http.DefaultClient.Do(req)
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
