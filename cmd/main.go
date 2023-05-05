package main

import (
	"fmt"
	"github.com/Krabik6/smart-contract-deployment/internal/compiler"
	"os"
)

func main() {

	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	compilers := compiler.NewCompiler(workDir, "ethereum/solc:stable", "artifacts")
	sourceCode := `// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract Bullshit {

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
	err = compilers.CompileBINSource(sourceCode)
	if err != nil {
		panic(err)
	}

	err = compilers.CompileABISource(sourceCode)
	if err != nil {
		panic(err)
	}

	//print abi
	bin, err := compilers.GetBytecode("Storage")
	if err != nil {
		panic(err)
	}
	println(bin)

	abi, err := compilers.GetJsonABI("Storage")
	if err != nil {
		panic(err)
	}
	fmt.Println(abi.Methods)

}

//docker run -v ${pwd}:/contract ethereum/solc:stable --abi --bin -o /contract/artifacts ./contract/Storage.sol --overwrite
