package deployer

import (
	"context"
	"github.com/ethereum/go-ethereum"
)

func (d *Deployer) EstimateGas(sourceCode string, optimize bool, runs int, args ...interface{}) (int, error) {
	auth := d.Ethereum.Auth     // auth is a pointer to a TransactOpts struct
	client := d.Ethereum.Client // client is a pointer to an ethclient.Client struct

	bytecode, err := d.GetBytecode(sourceCode, optimize, runs)
	if err != nil {
		return 0, err
	}

	estimateGas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     auth.From,
		To:       nil, // контракт еще не задеплоен
		GasPrice: auth.GasPrice,
		Value:    auth.Value,
		Data:     bytecode,
	})
	if err != nil {
		return 0, err
	}

	return int(estimateGas), nil
}
