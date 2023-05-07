package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// EthereumClient is a struct that contains the ethereum Client
type EthereumClient struct {
	Auth   *bind.TransactOpts
	Client *ethclient.Client
}

// NewEthereumClient creates a new ethereum Client
func NewEthereumClient(url string, privateKey string) (*EthereumClient, error) {
	client, err := NewClient(url)
	if err != nil {
		return nil, err
	}

	auth, err := NewAuth(privateKey, client)
	if err != nil {
		return nil, err
	}
	return &EthereumClient{
		Auth:   auth,
		Client: client,
	}, nil
}

func NewClient(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to create Ethereum Client: %v", err)
	}

	return client, nil
}

// Deploy wraps the bind.DeployContract function
func (c *EthereumClient) Deploy(bytecode []byte, abi abi.ABI, auth *bind.TransactOpts, client *ethclient.Client, args ...interface{}) (common.Address, *types.Transaction, interface{}, error) {

	contractAddress, tx, contract, err := bind.DeployContract(auth, abi, bytecode, client, args...)
	if err != nil {
		return contractAddress, tx, contract, err
	}

	return contractAddress, tx, contract, err
}

func NewAuth(privateKeyHex string, client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := ChainID(client)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("gasPrice: %v\n", gasPrice)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(1000000) // уменьшить значение газа до 1 миллиона

	return auth, nil
}
