package compiler

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"reflect"
)

func ConvertAndCheckArgs(args []interface{}, contractAbiJson *abi.ABI) ([]interface{}, error) {
	newArgs := make([]interface{}, len(args))
	for i := 0; i < len(args); i++ {
		log.Println(reflect.TypeOf(args[i]), "type arg ", i)
		abiArg := contractAbiJson.Constructor.Inputs[i]
		log.Println(abiArg.Type.String(), "constructor type arg ", i)

		switch abiArg.Type.T {
		case abi.UintTy:
			argUint, ok := args[i].(float64)
			if !ok {
				return nil, errors.New("failed to convert argument to uint")
			}
			newArgs[i] = big.NewInt(int64(argUint))
		case abi.StringTy:
			argStr, ok := args[i].(string)
			if !ok {
				return nil, errors.New("failed to convert argument to string")
			}
			newArgs[i] = argStr
		// добавьте другие типы данных здесь, если это необходимо
		default:
			return nil, errors.Errorf("unsupported argument type: %s", abiArg.Type.String())
		}
	}
	return newArgs, nil
}
