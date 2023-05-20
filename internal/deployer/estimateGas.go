package deployer

import (
	"context"
	"github.com/Krabik6/smart-contract-deployment/internal/eth"
	"github.com/ethereum/go-ethereum"
)

func (d *Deployer) EstimateGas(networkName string, network Network, bytecode []byte) (int, error) {
	_network, err := d.GetNetwork(networkName, network)
	if err != nil {
		return 0, err
	}
	ethereumClient, err := eth.NewEthereumClient(_network.Provider, _network.PrivateKey)
	if err != nil {
		return 0, err
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
	if err != nil {
		return 0, err
	}

	return int(estimateGas), nil
}
