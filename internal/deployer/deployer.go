package deployer

import (
	"context"
	"fmt"
	"github.com/Krabik6/smart-contract-deployment/internal/eth"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"reflect"
)

type Compiler interface {
	GetBytecode(sourceCode string, optimize bool, runs int) ([]byte, error)
	GetAbi(sourceCode string) (abi.ABI, error)
}

type Deployer struct {
	Ethereum *eth.EthereumClient
	Compiler
}

func NewDeployer(ethereum *eth.EthereumClient, compiler Compiler) *Deployer {
	return &Deployer{
		Ethereum: ethereum,
		Compiler: compiler,
	}
}

func (d *Deployer) Deploy(sourceCode string, optimize bool, runs int, args ...interface{}) (string, error) {
	// Deploy the contract
	address, tx, _, err := d.DeployWithArgs(sourceCode, optimize, runs, args...)
	if err != nil {
		return "", errors.Wrap(err, "failed to deploy contract")
	}

	fmt.Println("Contract address: ", address.Hex())
	fmt.Println("Transaction hash: ", tx.Hash().Hex())

	return address.Hex(), nil
}

func convertAndCheckArgs(args []interface{}, contractAbiJson *abi.ABI) ([]interface{}, error) {
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
			args[i] = big.NewInt(int64(argUint))
		case abi.StringTy:
			argStr, ok := args[i].(string)
			if !ok {
				return nil, errors.New("failed to convert argument to string")
			}
			args[i] = argStr
		// добавьте другие типы данных здесь, если это необходимо
		default:
			return nil, errors.Errorf("unsupported argument type: %s", abiArg.Type.String())
		}
	}
	return args, nil
}

func (d *Deployer) DeployWithArgs(sourceCode string, optimize bool, runs int, args ...interface{}) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	auth := d.Ethereum.Auth     // auth is a pointer to a TransactOpts struct
	client := d.Ethereum.Client // client is a pointer to an ethclient.Client struct
	// Load the contract GetAbi
	contractAbiJson, err := d.GetAbi(sourceCode)
	if err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "failed to load contract GetAbi")
	}
	// Load GetBytecode
	bytecode, err := d.GetBytecode(sourceCode, optimize, runs)
	if err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "failed to load bytecode")
	}
	estimateGas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     auth.From,
		To:       nil, // контракт еще не задеплоен
		GasPrice: auth.GasPrice,
		Value:    auth.Value,
		Data:     bytecode,
	})
	fmt.Println("Estimated gas: ", estimateGas)
	auth.GasLimit = uint64(float64(estimateGas) * 1.2)

	// Check if the contract expects arguments and set args to nil if none are expected
	//check type of args

	if len(args) == 1 && len(contractAbiJson.Constructor.Inputs) == 0 {
		args = nil
		log.Println("No arguments expected")
	}

	convertedArgs, err := convertAndCheckArgs(args, &contractAbiJson)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	// Deploy the contract
	address, tx, instance, err := bind.DeployContract(auth, contractAbiJson, bytecode, client, convertedArgs...)
	if err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "failed to deploy contract")
	}
	d.Ethereum.Auth.Nonce.Add(d.Ethereum.Auth.Nonce, big.NewInt(1))
	return address, tx, instance, nil
}
