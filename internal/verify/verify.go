package verify

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Verifier struct {
	ArgsEncoder ArgsEncoder
	Params      Params
}

func NewVerifier(argsEncoder ArgsEncoder) *Verifier {
	return &Verifier{
		ArgsEncoder: argsEncoder,
	}
}

func (v *Verifier) Verify(abi abi.ABI, params Params, constructorArguments ...interface{}) error {

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
