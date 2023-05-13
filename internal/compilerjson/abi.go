package compilerjson

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
	"log"
	"os/exec"
)

func (c *Compiler) GetAbi(sourceCode string) (abi.ABI, error) {
	cmd := exec.Command("docker", "run", "-i", "--rm", "-v", fmt.Sprintf("%s:/source", c.WorkDir), c.Image, "--standard-json", "/source/input.json")

	var output bytes.Buffer
	cmd.Stdout = &output

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to create stderr pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return abi.ABI{}, fmt.Errorf("failed to start Solidity compiler: %v", err)
	}

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error scanning compiler output: %v\n", err)
		}
	}()

	if err := cmd.Wait(); err != nil {
		return abi.ABI{}, fmt.Errorf("failed to compile Solidity contract: %v", err)
	}

	outputString := output.String()
	if outputString == "" {
		return abi.ABI{}, errors.New("no output from the Solidity compiler")
	}

	//log.Println("Output from the compiler:", outputString)

	// parse output to json format in SolcOutput struct
	var solcOutput SolcOutput
	err = json.Unmarshal([]byte(outputString), &solcOutput)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to parse compiler output: %v", err)
	}

	// Now you can access the data in your struct
	// For example:
	//log.Println(solcOutput)
	abiInterface := solcOutput.Contracts["smart_contracts/smart.sol"]["PublicStorageFuck"].Abi
	log.Println(abiInterface)
	//
	abiString, err := interfaceToString(abiInterface)
	if err != nil {
		return abi.ABI{}, err
	}

	abi, err := c.GetAbiFromString(abiString)

	return abi, nil
}

// convert abi format string to abi.ABI
func (c *Compiler) GetAbiFromString(abiString string) (abi.ABI, error) {
	abiBytes := []byte(abiString)

	// декодируем GetAbi из байтов
	abiJSON, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to decode GetAbi: %v", err)
	}

	return abiJSON, nil
}

func interfaceToString(value interface{}) (string, error) {
	// Check if the value is actually a string
	if str, ok := value.(string); ok {
		return str, nil
	}

	// If it's not a string, convert it using the fmt package
	return fmt.Sprintf("%v", value), nil
}
