package compiler

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"os/exec"
	"strings"
)

func (c *Compiler) GetBytecode(sourceCode string) ([]byte, error) {
	if c.WorkDir == "" {
		return nil, fmt.Errorf("work directory is not set")
	}

	fmt.Printf("Compiling Solidity contract in %s\n", c.WorkDir)

	cmd := exec.Command("docker", "run", "-i", "-v", fmt.Sprintf("%s:/contract", c.WorkDir), "ethereum/solc:stable", "--bin", "-", "--overwrite")
	cmd.Stdin = strings.NewReader(sourceCode)

	var output bytes.Buffer
	cmd.Stdout = &output

	// выводим процесс компиляции в консоль
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start Solidity compiler: %v", err)
	}
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			//fmt.Println(scanner.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("failed to compile Solidity contract: %v", err)
	}

	fmt.Println(output.String())

	// Remove "======= <stdin>:Bullshit =======\n" from the output
	//outputString := output.String()
	//abiStartIndex := strings.Index(outputString, "Binary:")
	//if abiStartIndex == -1 {
	//	return nil, fmt.Errorf("failed to find bytecode in output: %v", outputString)
	//}
	//binString := outputString[abiStartIndex+7:]
	//fmt.Println("binBytes", binString, "binString")
	binBytes := common.FromHex("608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100a1565b60405180910390f35b610073600480360381019061006e91906100ed565b61007e565b005b60008054905090565b8060008190555050565b6000819050919050565b61009b81610088565b82525050565b60006020820190506100b66000830184610092565b92915050565b600080fd5b6100ca81610088565b81146100d557600080fd5b50565b6000813590506100e7816100c1565b92915050565b600060208284031215610103576101026100bc565b5b6000610111848285016100d8565b9150509291505056fea26469706673582212205b7f104addc83c4319b1d441fb9ca06c8215d8f87dc25b55ba2c152cdfebd27d64736f6c63430008130033")

	fmt.Println("binBytes", binBytes, "binBytes")
	fmt.Println("608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100d9565b60405180910390f35b610073600480360381019061006e919061009d565b61007e565b005b60008054905090565b8060008190555050565b60008135905061009781610103565b92915050565b6000602082840312156100b3576100b26100fe565b5b60006100c184828501610088565b91505092915050565b6100d3816100f4565b82525050565b60006020820190506100ee60008301846100ca565b92915050565b6000819050919050565b600080fd5b61010c816100f4565b811461011757600080fd5b5056fea26469706673582212205071f5a59a07c341bf52e6ecc6c46eef83238e89dc9fe82b1736898b9f67815564736f6c63430008070033")
	binBytes = bytes.TrimSpace(binBytes)
	return binBytes, nil
}
