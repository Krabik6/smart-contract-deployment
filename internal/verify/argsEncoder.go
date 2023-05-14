package verify

import "github.com/ethereum/go-ethereum/accounts/abi"

type ArgsEncoder interface {
	EncodeConstructorArgs(abi abi.ABI, args ...interface{}) ([]byte, error)
}
