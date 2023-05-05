package compiler

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func (c *Compiler) Bytecode(sourceCode string) ([]byte, error) {
	//compile bytecode, get output
	c.CompileBINSource(sourceCode)
	return nil, nil
}

func (c *Compiler) GetBytecode(file string) ([]byte, error) {
	bytecodeFile := c.BytecodeFilename(file)
	bytecodeBytes, err := ioutil.ReadFile(bytecodeFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read bytecode file: %v", err)
	}
	return bytecodeBytes, nil
}

func (c *Compiler) CompileBINSource(sourceCode string) error {
	if c.WorkDir == "" {
		return fmt.Errorf("work directory is not set")
	}
	if c.OutputDir == "" {
		return fmt.Errorf("output directory is not set")
	}
	fmt.Printf("Compiling Solidity contract in %s\n", c.WorkDir)
	cmd := exec.Command("docker", "run", "-i", "-v", fmt.Sprintf("%s:/contract", c.WorkDir), "ethereum/solc:stable", "--bin", "-o", "/contract/"+c.OutputDir, "-", "--overwrite")
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

func (c *Compiler) BytecodeFilename(file string) string {
	return fmt.Sprintf("%s/%s.bin", c.OutputDir, file)
}
