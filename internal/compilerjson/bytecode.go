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

func (c *Compiler) GetBytecode(sourceCode string, inputPath string) ([]byte, error) {
	cmd := exec.Command("docker", "run", "-i", "--rm", "-v", fmt.Sprintf("%s:/source", c.WorkDir), "ethereum/solc:0.8.19", "--standard-json", "/source/input.json")

	var output bytes.Buffer
	cmd.Stdout = &output

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stderr pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start Solidity compiler: %v", err)
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
		return nil, fmt.Errorf("failed to compile Solidity contract: %v", err)
	}

	outputString := output.String()
	if outputString == "" {
		return nil, errors.New("no output from the Solidity compiler")
	}

	log.Println("Output from the compiler:", outputString)

	// parse output to json format in SolcOutput struct
	var solcOutput SolcOutput
	err = json.Unmarshal([]byte(outputString), &solcOutput)
	if err != nil {
		return nil, fmt.Errorf("failed to parse compiler output: %v", err)
	}

	// Now you can access the data in your struct
	// For example:
	//log.Println(solcOutput)
	bytecodeString := solcOutput.Contracts["smart_contracts/smart.sol"]["PublicStorageFuck"].Evm.Bytecode.Object
	binBytes := common.FromHex(bytecodeString)

	return binBytes, nil
}
