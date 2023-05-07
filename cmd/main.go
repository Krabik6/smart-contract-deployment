package main

import (
	"fmt"
	"github.com/Krabik6/smart-contract-deployment/internal/compiler"
	"github.com/Krabik6/smart-contract-deployment/internal/config"
	"github.com/Krabik6/smart-contract-deployment/internal/deployer"
	"github.com/Krabik6/smart-contract-deployment/internal/ethereum"
	"os"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
	eth, err := ethereum.NewEthereumClient(cfg.EnvConfig.Url, cfg.EnvConfig.PrivateKey)
	if err != nil {
		panic(err)
	}
	compilers := compiler.NewCompiler(workDir, "ethereum/solc:v0.8.7")
	deployers := deployer.NewDeployer(eth, compilers)

	sourceCode := `// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract Storage {

    uint256 number;

    /**
     * @dev Store value in variable
     * @param num value to store
     */
    function store(uint256 num) public {
        number = num;
    }

    /**
     * @dev Return value 
     * @return value of 'number'
     */
    function retrieve() public view returns (uint256){
        return number;
    }
}`

	addr, tx, err := deployers.Deploy(sourceCode)
	if err != nil {
		panic(err)
	}
	fmt.Println(addr.Hex())
	fmt.Println(tx.Hash().Hex())

}

//docker run -v ${pwd}:/contract ethereum/solc:stable --abi --bin -o /contract/artifacts ./contract/Storage.sol --overwrite
