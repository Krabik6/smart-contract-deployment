package main

import (
	"github.com/Krabik6/smart-contract-deployment/internal/compilerjson"
	"log"
	"os"
)

func main() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	compilersjson := compilerjson.NewCompiler(workDir, "ethereum/solc:0.8.19")
	//abi, err := compilersjson.GetAbi("./smart_contracts/smart.sol")
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(abi)

	bytecode, err := compilersjson.GetBytecode("", "input.json")
	if err != nil {
		panic(err)
	}
	log.Println(bytecode)
	//mainSolPath := "./smart_contracts/smart.sol"
	//outputPath := "input.json"
	//
	//c := inputgenerator.NewCompiler()
	//err := c.WriteJSONInput(mainSolPath, outputPath, true, 200)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//log.Println("Creating input.json file...")
	//
	////wait 50 seconds and panic
	//time.Sleep(50 * time.Second)
	//panic("panic")

	//cfg, err := config.Load()
	//if err != nil {
	//	panic(err)
	//}
	//
	//workDir, err := os.Getwd()
	//if err != nil {
	//	panic(err)
	//}
	//eth, err := eth.NewEthereumClient(cfg.EnvConfig.Url, cfg.EnvConfig.PrivateKey)
	//if err != nil {
	//	panic(err)
	//}
	//compilers := compiler.NewCompiler(workDir, cfg.AppConfig.Image)
	//deployers := deployer.NewDeployer(eth, compilers)
	//verifiers := verify.NewVerifier(compilers)
	//handlers := handler.NewHandler(deployers, compilers, verifiers)
	//
	//srv := apiserver.NewServer()
	//
	//if err := srv.Run(cfg.Server.Port, handlers.InitRouts()); err != nil {
	//	panic(err)
	//}

	//fmt.Println(handlers)
	//	sourceCode := `// SPDX-License-Identifier: GPL-3.0
	//
	//pragma solidity >=0.7.0 <0.9.0;
	//
	///**
	// * @title Storage
	// * @dev Store & retrieve value in a variable
	// * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
	// */
	//contract Storage {
	//
	//    uint256 number;
	//
	//    /**
	//     * @dev Store value in variable
	//     * @param num value to store
	//     */
	//    function store(uint256 num) public {
	//        number = num;
	//    }
	//
	//    /**
	//     * @dev Return value
	//     * @return value of 'number'
	//     */
	//    function retrieve() public view returns (uint256){
	//        return number;
	//    }
	//}`

	//addr, err := deployers.Deploy(sourceCode)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(addr)

}
