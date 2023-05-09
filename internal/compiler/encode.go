package compiler

import (
	"fmt"
)

func (c *Compiler) EncodeConstructorArgs(sourceCode string, args ...interface{}) ([]byte, error) {
	abiJSON, err := c.GetAbi(sourceCode)
	parsedArgs, err := c.ConvertAndCheckArgs(args, &abiJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to convert and check constructor arguments: %v", err)
	}

	encodedArgs, err := abiJSON.Pack("", parsedArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to encode constructor arguments: %v", err)
	}
	return encodedArgs, nil
}

func (c *Compiler) EncodeFunctionCall(sourceCode string, functionName string, args ...interface{}) ([]byte, error) {
	abiJSON, err := c.GetAbi(sourceCode)
	parsedArgs, err := c.ConvertAndCheckArgs(args, &abiJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to convert and check constructor arguments: %v", err)
	}

	encodedArgs, err := abiJSON.Pack(functionName, parsedArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to encode function arguments: %v", err)
	}
	return encodedArgs, nil
}
