package compiler

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io/ioutil"
	"os/exec"
	"strings"
)

func (c *Compiler) GetJsonABI(file string) (abi.ABI, error) {
	abiBytes, err := c.ByteABI(file)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to get ABI: %v", err)
	}
	abiJSON, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to get ABI: %v", err)
	}
	return abiJSON, nil

}

func (c *Compiler) ByteABI(file string) ([]byte, error) {
	abiFile := c.AbiFilename(file)
	abiBytes, err := ioutil.ReadFile(abiFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read ABI file: %v", err)
	}
	return abiBytes, nil
}

func (c *Compiler) CompileABISource(sourceCode string) error {
	if c.WorkDir == "" {
		return fmt.Errorf("work directory is not set")
	}
	if c.OutputDir == "" {
		return fmt.Errorf("output directory is not set")
	}
	fmt.Printf("Compiling Solidity contract in %s\n", c.WorkDir)
	cmd := exec.Command("docker", "run", "-i", "-v", fmt.Sprintf("%s:/contract", c.WorkDir), "ethereum/solc:stable", "--abi", "-o", "/contract/"+c.OutputDir, "-", "--overwrite")
	cmd.Stdin = strings.NewReader(sourceCode)

	// выводим процесс компиляции в консоль
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start Solidity compiler: %v", err)
	}
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to compile Solidity contract: %v", err)
	}

	fmt.Println("Compiled successfully")
	return nil
}

func (c *Compiler) AbiFilename(file string) string {
	return fmt.Sprintf("%s/%s.abi", c.OutputDir, file)
}
