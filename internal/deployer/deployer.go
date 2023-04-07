package deployer

import (
	"github.com/Krabik6/smart-contract-deployment/internal/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type Compiler interface {
	Bytecode(sourceCode string) ([]byte, error)
	ABI(sourceCode string) (abi.ABI, error)
}

type Deployer struct {
	ethereum *ethereum.EthereumClient
	Compiler
}

func NewDeployer(ethereum *ethereum.EthereumClient, compiler Compiler) *Deployer {
	return &Deployer{
		ethereum: ethereum,
		Compiler: compiler,
	}
}

func (d *Deployer) Deploy(sourceCode string, args ...interface{}) (common.Address, *types.Transaction, error) {
	auth := d.ethereum.Auth     // auth is a pointer to a TransactOpts struct
	client := d.ethereum.Client // client is a pointer to an ethclient.Client struct
	// Load the contract ABI
	contractAbiJson, err := d.ABI(sourceCode)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "failed to load contract ABI")
	}

	// Load Bytecode
	bytecode, err := d.Bytecode(sourceCode)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "failed to load bytecode")
	}

	// Deploy the contract
	address, tx, _, err := bind.DeployContract(auth, contractAbiJson, bytecode, client, args...)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "failed to deploy contract")
	}

	return address, tx, nil
}

//// deploy by source code + arg
//func (c *EthereumClient) Deploy(sourceCode string, args []interface{}, value int) (string, error) {
//	c.auth.Value = big.NewInt(int64(value))
//	// Load the contract ABI
//	const contractAbiJson = `[...]`
//	contractAbi, err := abi.JSON(strings.NewReader(contractAbiJson))
//	if err != nil {
//		log.Fatal(err)
//	}
//	contractAddress, tx, _, err := bind.DeployContract(c.auth, contractAbi, []byte(sourceCode), c.client, args...)
//
//	if err != nil {
//		return "", err
//	}
//
//	_, err = bind.WaitMined(context.Background(), c.client, tx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("Contract address:", contractAddress.Hex(), tx.Hash().Hex())
//	return contractAddress.Hex(), nil
//}
