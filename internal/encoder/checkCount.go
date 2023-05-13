package encoder

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
)

func checkArgumentCount(args []interface{}, abiInputs []abi.Argument) error {
	if len(args) != len(abiInputs) {
		return errors.New(fmt.Sprintf("wrong number of arguments: %d given, %d expected", len(args), len(abiInputs)))
	}
	return nil
}
