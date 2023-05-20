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

//type Compiler interface {
//	GetBytecode(sourceCode string, optimize bool, runs int) ([]byte, error)
//	GetAbi(sourceCode string) (abi.ABI, error)
//}

type Deployer struct {
	//Ethereum *eth.EthereumClient
	Networks map[string]Network
}

type Network struct {
	Provider   string
	PrivateKey string
}

func NewDeployer(networks map[string]Network) *Deployer {
	return &Deployer{
		//Ethereum: ethereum,
		Networks: networks,
	}
}

func (d *Deployer) Deploy(networkName string, network Network, bytecode []byte, abi abi.ABI, args ...interface{}) (string, error) {
	// Deploy the contract
	address, tx, _, err := d.DeployWithArgs(networkName, network, bytecode, abi, args...)
	if err != nil {
		return "", errors.Wrap(err, "failed to deploy contract")
	}

	fmt.Println("Contract address: ", address.Hex())
	fmt.Println("Transaction hash: ", tx.Hash().Hex())

	return address.Hex(), nil
}

// todo interface for convert and check args
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

func (d *Deployer) DeployWithArgs(networkName string, network Network, bytecode []byte, abi abi.ABI, args ...interface{}) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	_network, err := d.GetNetwork(networkName, network)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	ethereumClient, err := eth.NewEthereumClient(_network.Provider, _network.PrivateKey)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	auth := ethereumClient.Auth     // auth is a pointer to a TransactOpts struct
	client := ethereumClient.Client // client is a pointer to an ethclient.Client struct

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
	if len(args) == 1 && len(abi.Constructor.Inputs) == 0 {
		args = nil
		log.Println("No arguments expected")
	}

	convertedArgs, err := convertAndCheckArgs(args, &abi)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	//log.Print(bytecode)
	//log.Print(abi)

	// Deploy the contract
	address, tx, instance, err := bind.DeployContract(auth, abi, bytecode, client, convertedArgs...)
	if err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "failed to deploy contract")
	}
	ethereumClient.Auth.Nonce.Add(ethereumClient.Auth.Nonce, big.NewInt(1))
	return address, tx, instance, nil
}
