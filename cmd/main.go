package main

import (
	"github.com/Krabik6/smart-contract-deployment/internal/apiserver"
	"github.com/Krabik6/smart-contract-deployment/internal/compiler"
	"github.com/Krabik6/smart-contract-deployment/internal/compilerjson"
	"github.com/Krabik6/smart-contract-deployment/internal/config"
	"github.com/Krabik6/smart-contract-deployment/internal/deployer"
	"github.com/Krabik6/smart-contract-deployment/internal/encoder"
	"github.com/Krabik6/smart-contract-deployment/internal/eth"
	"github.com/Krabik6/smart-contract-deployment/internal/handler"
	"github.com/Krabik6/smart-contract-deployment/internal/inputgenerator"
	"github.com/Krabik6/smart-contract-deployment/internal/verify"
	"os"
)

func main() {

	//bytecode, err := compilersjson.GetBytecode("", "input.json")
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(bytecode)

	//abi, err := compilersjson.GetAbi("")
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(abi)

	//_mainSolPath := "smart_contracts/smart.sol"
	//_mainSolName := "PublicStorageFuck"
	//
	//c := inputgenerator.NewCompiler()
	//jsonInput, mainSolPath, err := c.GenerateJSONInput(_mainSolPath, true, 200)
	//if err != nil {
	//	panic(err)
	//}
	//
	//workDir, err := os.Getwd()
	//if err != nil {
	//	panic(err)
	//}
	//
	//compilersjson := compilerjson.NewCompiler(workDir, "ethereum/solc:0.8.19")
	//bytecode, err := compilersjson.GetBytecode(jsonInput, mainSolPath, _mainSolName)
	//if err != nil {
	//	panic(err)
	//}
	//log.Print(bytecode)
	//
	//abi, err := compilersjson.GetAbi(jsonInput, mainSolPath, _mainSolName)
	//if err != nil {
	//	panic(err)
	//}
	//log.Println("abi: ", abi)
	//
	////wait 50 seconds and panic
	//time.Sleep(50 * time.Second)
	//
	//panic("panic")

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	eth, err := eth.NewEthereumClient(cfg.EnvConfig.Url, cfg.EnvConfig.PrivateKey)
	if err != nil {
		panic(err)
	}
	inputGenerators := inputgenerator.NewInputGenerator()
	compilersjson := compilerjson.NewCompiler(workDir, "ethereum/solc:0.8.19")
	argsEncoders := encoder.NewEncoder()
	compilers := compiler.NewCompiler(workDir, cfg.AppConfig.Image)
	deployers := deployer.NewDeployer(eth)
	verifiers := verify.NewVerifier(argsEncoders)
	handlers := handler.NewHandler(deployers, compilers, verifiers, argsEncoders, compilersjson, inputGenerators)

	srv := apiserver.NewServer()

	if err := srv.Run(cfg.Server.Port, handlers.InitRouts()); err != nil {
		panic(err)
	}

}
