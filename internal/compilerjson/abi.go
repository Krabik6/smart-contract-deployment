package compilerjson

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
	"os/exec"
	"strings"
)

func (c *Compiler) GetAbi(sourceCode string) (abi.ABI, error) {
	if c.WorkDir == "" {
		return abi.ABI{}, fmt.Errorf("work directory is not set")
	}

	cmd := exec.Command("docker", "run", "-i", "-v", fmt.Sprintf("%s:/contract", c.WorkDir), c.Image, "--standard-json", "/contract/input.json")
	cmd.Dir = c.WorkDir

	var output bytes.Buffer
	cmd.Stdout = &output

	// выводим процесс компиляции в консоль
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return abi.ABI{}, fmt.Errorf("failed to start Solidity compiler: %v", err)
	}
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return abi.ABI{}, fmt.Errorf("failed to compile Solidity contract: %v", err)
	}

	outputString := output.String()
	abiStartIndex := strings.Index(outputString, "[{")
	if abiStartIndex == -1 {
		return abi.ABI{}, fmt.Errorf("failed to find ABI in output: %v", outputString)
	}
	abiString := outputString[abiStartIndex:]
	log.Println(abiString)

	abiJSON, err := c.GetAbiFromString(abiString)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to decode ABI: %v", err)
	}

	return abiJSON, nil
}

// GetAbiFromString convert ABI format string to abi.ABI
func (c *Compiler) GetAbiFromString(abiString string) (abi.ABI, error) {
	abiBytes := []byte(abiString)

	// декодируем ABI из байтов
	abiJSON, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to decode ABI: %v", err)
	}

	return abiJSON, nil
}
