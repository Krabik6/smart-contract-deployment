package compiler

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"os/exec"
	"strings"
)

func (c *Compiler) GetBytecode(sourceCode string, optimize bool, runs int) ([]byte, error) {
	if c.WorkDir == "" {
		return nil, errors.New("work directory is not set")
	}

	var optimizeFlags []string
	if optimize {
		fmt.Println("Enabling optimization with", runs, "runs")
		optimizeFlags = []string{"--optimize", fmt.Sprintf("--optimize-runs=%d", runs)}
	}

	cmd := exec.Command("docker", "run", "-i", "-v", fmt.Sprintf("%s:/contract", c.WorkDir))
	cmd.Args = append(cmd.Args, c.Image, "--bin", "-", "--overwrite")
	if optimize {
		cmd.Args = append(cmd.Args, optimizeFlags...)
	}
	cmd.Stdin = strings.NewReader(sourceCode)

	var output bytes.Buffer
	cmd.Stdout = &output

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stderr pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start Solidity compiler: %v", err)
	}

	// Read and handle compiler stderr output in a separate goroutine
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

	// Ищем индекс строки "Binary:"
	binIndex := strings.Index(outputString, "Binary:")
	if binIndex == -1 {
		return nil, errors.New("failed to find bytecode in the output")
	}

	// Извлекаем байткод после "Binary:"
	binString := strings.TrimSpace(outputString[binIndex+len("Binary:"):])

	binBytes := common.FromHex(binString)

	if len(binBytes) == 0 {
		fmt.Printf("Output string from the compiler: %s\n", outputString)
		return nil, errors.New("bytecode is empty")
	}

	return binBytes, nil
}
