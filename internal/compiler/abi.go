package compiler

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"os/exec"
	"strings"
)

func (c *Compiler) GetABI(sourceCode string) (abi.ABI, error) {
	if c.WorkDir == "" {
		return abi.ABI{}, fmt.Errorf("work directory is not set")
	}

	fmt.Printf("Compiling Solidity contract in %s\n", c.WorkDir)

	cmd := exec.Command("docker", "run", "-i", "-v", fmt.Sprintf("%s:/contract", c.WorkDir), "ethereum/solc:stable", "--abi", "-", "--overwrite")
	cmd.Stdin = strings.NewReader(sourceCode)

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

	// Remove "======= <stdin>:Bullshit =======\n" from the output
	outputString := output.String()
	abiStartIndex := strings.Index(outputString, "[{")
	if abiStartIndex == -1 {
		return abi.ABI{}, fmt.Errorf("failed to find GetABI in output: %v", outputString)
	}
	abiString := outputString[abiStartIndex:]
	abiBytes := []byte(abiString)

	// декодируем GetABI из байтов
	abiJSON, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to decode GetABI: %v", err)
	}
	return abiJSON, nil
}
