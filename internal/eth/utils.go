package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func ChainID(client *ethclient.Client) (*big.Int, error) {
	return client.ChainID(context.Background())
}
