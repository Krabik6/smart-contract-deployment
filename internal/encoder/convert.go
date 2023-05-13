package encoder

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"reflect"
)

func ConvertArgs(args []interface{}, contractAbiJson *abi.ABI) ([]interface{}, error) {
	newArgs := make([]interface{}, len(args))
	err := checkArgumentCount(args, contractAbiJson.Constructor.Inputs)
	if err != nil {
		return nil, err
	}

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
		case abi.BoolTy:
			argBool, ok := args[i].(bool)
			if !ok {
				return nil, errors.New("failed to convert argument to bool")
			}
			newArgs[i] = argBool
		case abi.AddressTy:
			argAddress, ok := args[i].(common.Address)
			if !ok {
				return nil, errors.New("failed to convert argument to address")
			}
			newArgs[i] = argAddress
		case abi.BytesTy:
			argBytes, ok := args[i].([]byte)
			if !ok {
				return nil, errors.New("failed to convert argument to bytes")
			}
			newArgs[i] = argBytes
		case abi.IntTy:
			argInt, ok := args[i].(float64)
			if !ok {
				return nil, errors.New("failed to convert argument to int")
			}
			newArgs[i] = big.NewInt(int64(argInt))
		case abi.FixedBytesTy:
			argFixedBytes, ok := args[i].([]byte)
			if !ok {
				return nil, errors.New("failed to convert argument to fixed bytes")
			}
			newArgs[i] = argFixedBytes
		case abi.FixedPointTy:
			argFixedPoint, ok := args[i].(float64)
			if !ok {
				return nil, errors.New("failed to convert argument to fixed point")
			}
			newArgs[i] = big.NewRat(int64(argFixedPoint), 1)
		default:
			return nil, errors.Errorf("unsupported argument type: %s", abiArg.Type.String())
		}
	}
	return newArgs, nil
}
