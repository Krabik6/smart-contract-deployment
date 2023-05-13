package encoder

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Encoder struct {
}

func NewEncoder() *Encoder {
	return &Encoder{}
}

func (e *Encoder) EncodeConstructorArgs(abi abi.ABI, args ...interface{}) ([]byte, error) {
	parsedArgs, err := ConvertArgs(args, &abi)
	if err != nil {
		return nil, fmt.Errorf("failed to convert and check constructor arguments: %v", err)
	}

	encodedArgs, err := abi.Pack("", parsedArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to encode constructor arguments: %v", err)
	}
	return encodedArgs, nil
}

func (e *Encoder) EncodeFunctionCall(abi abi.ABI, functionName string, args ...interface{}) ([]byte, error) {
	parsedArgs, err := ConvertArgs(args, &abi)
	if err != nil {
		return nil, fmt.Errorf("failed to convert and check constructor arguments: %v", err)
	}

	encodedArgs, err := abi.Pack(functionName, parsedArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to encode function arguments: %v", err)
	}
	return encodedArgs, nil
}
