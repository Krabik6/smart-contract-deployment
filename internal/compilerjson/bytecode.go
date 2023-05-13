package compilerjson

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"log"
	"os/exec"
)

func (c *Compiler) GetBytecode(inputJSON []byte, contractPath, contractName string) ([]byte, error) {
	cmd := exec.Command("docker", "run", "-i", "--rm", "-v", fmt.Sprintf("%s:/source", c.WorkDir), c.Image, "--standard-json")

	var output bytes.Buffer
	cmd.Stdout = &output

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stderr pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start Solidity compiler: %v", err)
	}

	go func() {
		defer stdin.Close()
		_, _ = stdin.Write(inputJSON)
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
		return nil, fmt.Errorf("failed to compile Solidity contract: %v", err)
	}

	outputString := output.String()
	if outputString == "" {
		return nil, errors.New("no output from the Solidity compiler")
	}

	// parse output to json format in SolcOutput struct
	var solcOutput SolcOutput
	err = json.Unmarshal([]byte(outputString), &solcOutput)
	if err != nil {
		return nil, fmt.Errorf("failed to parse compiler output: %v", err)
	}

	log.Println(solcOutput.Contracts[contractPath][contractName], "1")

	// Access the ABI of the first contract
	bytecodeInterface := solcOutput.Contracts[contractPath][contractName].Evm.Bytecode.Object
	bytecodeBytes := common.FromHex(bytecodeInterface)

	return bytecodeBytes, nil
}
