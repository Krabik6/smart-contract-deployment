package compilerjson

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func (c *Compiler) GetBytecode(sourceCode string, inputPath string) ([]byte, error) {
	cmd := exec.Command("docker", "run", "-i", "--rm", "-v", fmt.Sprintf("%s:/source", c.WorkDir), "ethereum/solc:0.8.19", "--standard-json", "/source/input.json")
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

	// Ищем байткод с помощью регулярного выражения
	re := regexp.MustCompile(`"object":"\s*([0-9a-fA-F]+)`)
	matches := re.FindStringSubmatch(outputString)
	if len(matches) < 2 {
		return nil, errors.New("failed to find bytecode in the output")
	}

	binString := strings.TrimSpace(matches[1])

	binBytes := common.FromHex(binString)

	if len(binBytes) == 0 {
		fmt.Printf("Output string from the compiler: %s\n", outputString)
		return nil, errors.New("bytecode is empty")
	}

	// binBytes to string
	log.Println("Bytecode:", hex.EncodeToString(binBytes))

	return binBytes, nil
}
