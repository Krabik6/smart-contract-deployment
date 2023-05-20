package verify

import "github.com/pkg/errors"

func (v *Verifier) GetNetwork(networkName string, network Network) (Network, error) {
	// if networkName is empty and network.Apikey or network.Url is empty
	if networkName == "" && (network.Apikey == "" || network.Url == "") {
		return Network{}, errors.New("network name is empty and network apikey or url is empty")
	}

	// return default network with optional custom api key
	if networkName != "" {
		_network, err := v.getDefaultNetwork(networkName)
		if err != nil {
			return Network{}, err
		}
		apiKey, ok := v.getCustomApiKey(network)
		if ok {
			_network.Apikey = apiKey
		}
		return _network, nil
	}

	// return custom network
	return v.getCustomNetwork(network)
}

func (v *Verifier) getDefaultNetwork(networkName string) (Network, error) {
	network := Network{}
	if networkName == "" {
		return network, errors.New("network name is empty")
	}
	if v.Networks == nil {
		return network, errors.New("networks map is empty")
	}
	network, ok := v.Networks[networkName]
	if !ok {
		return network, errors.New("network not found")
	}
	if network.Apikey == "" {
		return network, errors.New("network apikey is empty")
	}
	if network.Url == "" {
		return network, errors.New("network url is empty")
	}
	return network, nil
}

// get custom api key for way where user want to use our default network, but with custom api key
func (v *Verifier) getCustomApiKey(network Network) (string, bool) {
	if network.Apikey == "" {
		return "", false
	}
	return network.Apikey, true
}

func (v *Verifier) getCustomNetwork(network Network) (Network, error) {
	if network.Apikey == "" {
		return Network{}, errors.New("network apikey is empty")
	}
	if network.Url == "" {
		return Network{}, errors.New("network url is empty")
	}
	return network, nil
}
