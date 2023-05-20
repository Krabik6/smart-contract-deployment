package deployer

import "github.com/pkg/errors"

func (d *Deployer) GetNetwork(networkName string, network Network) (Network, error) {
	// if networkName is empty and network.Apikey or network.Url is empty
	if networkName == "" && (network.Provider == "" || network.PrivateKey == "") {
		return Network{}, errors.New("network name is empty and network private key or provider is empty")
	}

	// return default network with optional custom api key
	if networkName != "" {
		_network, err := d.getDefaultNetwork(networkName)
		if err != nil {
			return Network{}, err
		}
		apiKey, ok := d.getCustomApiKey(network)
		if ok {
			_network.PrivateKey = apiKey
		}
		return _network, nil
	}

	// return custom network
	return d.getCustomNetwork(network)
}

func (d *Deployer) getDefaultNetwork(networkName string) (Network, error) {
	network := Network{}
	if networkName == "" {
		return network, errors.New("network name is empty")
	}
	if d.Networks == nil {
		return network, errors.New("networks map is empty")
	}
	network, ok := d.Networks[networkName]
	if !ok {
		return network, errors.New("network not found")
	}
	if network.PrivateKey == "" {
		return network, errors.New("private key is empty")
	}
	if network.Provider == "" {
		return network, errors.New("provider url is empty")
	}
	return network, nil
}

// get custom api key for way where user want to use our default network, but with custom api key
func (d *Deployer) getCustomApiKey(network Network) (string, bool) {
	if network.PrivateKey == "" {
		return "", false
	}
	return network.PrivateKey, true
}

func (d *Deployer) getCustomNetwork(network Network) (Network, error) {
	if network.PrivateKey == "" {
		return Network{}, errors.New("network apikey is empty")
	}
	if network.Provider == "" {
		return Network{}, errors.New("network url is empty")
	}
	return network, nil
}
