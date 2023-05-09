package compiler

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func EncodeConstructorArgs(abiJSON abi.ABI, args ...interface{}) ([]byte, error) {
	encodedArgs, err := abiJSON.Pack("", args...)
	if err != nil {
		return nil, fmt.Errorf("failed to encode constructor arguments: %v", err)
	}
	return encodedArgs, nil
}

func EncodeFunctionCall(abiJSON abi.ABI, functionName string, args ...interface{}) ([]byte, error) {
	encodedArgs, err := abiJSON.Pack(functionName, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to encode function arguments: %v", err)
	}
	return encodedArgs, nil
}
