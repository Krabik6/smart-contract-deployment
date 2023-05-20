package verify

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
)

type Verifier struct {
	ArgsEncoder ArgsEncoder
	Params      Params
	Networks    map[string]Network
}

type Network struct {
	Apikey string
	Url    string
}

func NewVerifier(argsEncoder ArgsEncoder, network map[string]Network) *Verifier {
	return &Verifier{
		ArgsEncoder: argsEncoder,
		Networks:    network,
	}
}

func (v *Verifier) Verify(networkName string, network Network, abi abi.ABI, params Params, constructorArguments ...interface{}) error {

	v.logParams(params)

	_network, err := v.GetNetwork(networkName, network)
	if err != nil {
		return err
	}
	log.Println("Network:", _network)

	_params, err := v.prepareParams(_network.Apikey, abi, params, constructorArguments...)
	if err != nil {
		return err
	}

	req, err := v.prepareRequest(_network.Url, _params)
	if err != nil {
		return err
	}

	err = v.sendRequestAndParseResponse(req)
	if err != nil {
		return err
	}

	return nil
}
