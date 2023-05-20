package main

import (
	"github.com/Krabik6/smart-contract-deployment/internal/apiserver"
	"github.com/Krabik6/smart-contract-deployment/internal/compiler"
	"github.com/Krabik6/smart-contract-deployment/internal/compilerjson"
	"github.com/Krabik6/smart-contract-deployment/internal/config"
	"github.com/Krabik6/smart-contract-deployment/internal/deployer"
	"github.com/Krabik6/smart-contract-deployment/internal/encoder"
	"github.com/Krabik6/smart-contract-deployment/internal/handler"
	"github.com/Krabik6/smart-contract-deployment/internal/inputgenerator"
	"github.com/Krabik6/smart-contract-deployment/internal/verify"
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

	networks := map[string]verify.Network{
		"mumbai": {
			Apikey: "IXQV2ZCWX4X3KZ8RDSHNYARAF8DR6F2DZ5",
			Url:    "https://api-testnet.polygonscan.com/api",
		},
		"mainnet": {
			Apikey: "PEKY2JR6FTUHD6MJDVQGFTJDASBGSG6BSA",
			Url:    "https://api.etherscan.io/api",
		},
		"polygon": {
			Apikey: "IXQV2ZCWX4X3KZ8RDSHNYARAF8DR6F2DZ5",
			Url:    "https://api.polygonscan.com/api",
		},
		"goerli": {
			Apikey: "PEKY2JR6FTUHD6MJDVQGFTJDASBGSG6BSA",
			Url:    "https://api-goerli.etherscan.io/api",
		},
		"bsc-testnet": {
			Apikey: "T6KJBBV77UATD6TXPMH52TYGJVGVNH1EVW",
			Url:    "https://api-testnet.bscscan.com/api",
		},
		"sepolia": {
			Apikey: "HT1C3XJ7ZNMIDXBT32H1GMPN44P9P5A2EH",
			Url:    "https://api-sepolia.etherscan.io/api",
		},
	}
	deployNetworks := map[string]deployer.Network{
		"mumbai": {
			Provider:   cfg.MumbaiProvider,
			PrivateKey: cfg.MumbaiPrivateKey,
		},
		"bsc-testnet": {
			Provider:   cfg.BscTestnetProvider,
			PrivateKey: cfg.BscTestnetPrivateKey,
		},
		"sepolia": {
			Provider:   cfg.SepoliaProvider,
			PrivateKey: cfg.SepoliaPrivateKey,
		},
	}

	inputGenerators := inputgenerator.NewInputGenerator()
	compilersjson := compilerjson.NewCompiler(workDir, "ethereum/solc:0.8.19")
	argsEncoders := encoder.NewEncoder()
	compilers := compiler.NewCompiler(workDir, cfg.AppConfig.Image)
	deployers := deployer.NewDeployer(deployNetworks)
	verifiers := verify.NewVerifier(argsEncoders, networks)
	handlers := handler.NewHandler(deployers, compilers, verifiers, argsEncoders, compilersjson, inputGenerators)

	srv := apiserver.NewServer()

	if err := srv.Run(cfg.Server.Port, handlers.InitRouts()); err != nil {
		panic(err)
	}

}
