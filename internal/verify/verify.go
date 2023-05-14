package verify

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Verifier struct {
	ArgsEncoder ArgsEncoder
}

func NewVerifier(argsEncoder ArgsEncoder) *Verifier {
	return &Verifier{
		ArgsEncoder: argsEncoder,
	}
}

//
//// Compiler is the interface that wraps the Compile method.
//type Compiler interface {
//	EncodeConstructorArgs(sourceCode string, args ...interface{}) ([]byte, error)
//}
//
//type CompilerJson interface {
//	GetBytecode(inputJSON []byte, contractPath, contractName string) ([]byte, error)
//	GetAbi(inputJSON []byte, contractPath, contractName string) (abi.ABI, error)
//}

type ArgsEncoder interface {
	EncodeConstructorArgs(abi abi.ABI, args ...interface{}) ([]byte, error)
}

func (v *Verifier) Verify(abi abi.ABI, params Params, constructorArguments ...interface{}) error {
	err := v.validateParams(params)
	if err != nil {
		return err
	}

	v.logParams(params)

	_params, err := v.prepareParams(abi, params, constructorArguments...)
	if err != nil {
		return err
	}

	req, err := v.prepareRequest(_params)
	if err != nil {
		return err
	}

	err = v.sendRequestAndParseResponse(req)
	if err != nil {
		return err
	}

	return nil
}
