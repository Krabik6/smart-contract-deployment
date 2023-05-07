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
	GetBytecode(sourceCode string) ([]byte, error)
	GetABI(sourceCode string) (abi.ABI, error)
}

type Deployer struct {
	Ethereum *ethereum.EthereumClient
	Compiler
}

func NewDeployer(ethereum *ethereum.EthereumClient, compiler Compiler) *Deployer {
	return &Deployer{
		Ethereum: ethereum,
		Compiler: compiler,
	}
}

func (d *Deployer) Deploy(sourceCode string, args ...interface{}) (common.Address, *types.Transaction, error) {
	auth := d.Ethereum.Auth     // auth is a pointer to a TransactOpts struct
	client := d.Ethereum.Client // client is a pointer to an ethclient.Client struct
	// Load the contract GetABI
	contractAbiJson, err := d.GetABI(sourceCode)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "failed to load contract GetABI")
	}

	// Load GetBytecode
	bytecode, err := d.GetBytecode(sourceCode)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "failed to load bytecode")
	}

	// SwapperABI is the input ABI used to generate the binding from.
	//const SwapperABI = "[\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"retrieve\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"num\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"store\",\n\t\t\"outputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t}\n]"
	//// SwapperBin is the compiled bytecode used for deploying new contracts.
	//var SwapperBin = "608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100d9565b60405180910390f35b610073600480360381019061006e919061009d565b61007e565b005b60008054905090565b8060008190555050565b60008135905061009781610103565b92915050565b6000602082840312156100b3576100b26100fe565b5b60006100c184828501610088565b91505092915050565b6100d3816100f4565b82525050565b60006020820190506100ee60008301846100ca565b92915050565b6000819050919050565b600080fd5b61010c816100f4565b811461011757600080fd5b5056fea26469706673582212205071f5a59a07c341bf52e6ecc6c46eef83238e89dc9fe82b1736898b9f67815564736f6c63430008070033"

	//abiObj, err := abi.JSON(strings.NewReader(SwapperABI))
	//if err != nil {
	//	panic(err)
	//}
	//tx to transfer ether to the zero address

	// Deploy the contract
	address, tx, _, err := bind.DeployContract(auth, contractAbiJson, bytecode, client, args...)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "failed to deploy contract")
	}

	return address, tx, nil
}
