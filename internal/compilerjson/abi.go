package compilerjson

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
	"os/exec"
)

func (c *Compiler) GetAbi(inputJSON []byte) (abi.ABI, error) {
	cmd := exec.Command("docker", "run", "-i", "--rm", "-v", fmt.Sprintf("%s:/source", c.WorkDir), c.Image, "--standard-json")

	var output bytes.Buffer
	cmd.Stdout = &output

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to create stdin pipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to create stderr pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return abi.ABI{}, fmt.Errorf("failed to start Solidity compiler: %v", err)
	}

	go func() {
		defer stdin.Close()
		_, err = stdin.Write(inputJSON)
	}()

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

	// Преобразование abiInterface в []byte
	abiBytes, err := json.Marshal(abiInterface)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to convert abiInterface to bytes: %v", err)
	}

	// Декодирование ABI из []byte
	abiJSON, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to decode ABI: %v", err)
	}

	return abiJSON, nil
}
