package handler

import "github.com/gin-gonic/gin"

/*
- /deploy (POST) [source code, constructor arguments] ⇒ [address]
- /abi (GET) [source code] ⇒ [abi]
- /bytecode (GET) [source code] ⇒ [bytecode]
- /abiBytecode [source code] ⇒ [abi, bytecode]
*/

//api endpoints like /deploy, /abi, /bytecode, /abiBytecode using gin framework

type Handler struct {
	Deployer       Deployer
	Compiler       Compiler
	Verifier       Verifier
	ArgsEncoder    ArgsEncoder
	CompilerJson   CompilerJson
	InputGenerator InputGenerator
}

func NewHandler(deployer Deployer, compiler Compiler, verifier Verifier, argsEncoder ArgsEncoder, compilerJson CompilerJson, inputGenerator InputGenerator) *Handler {
	return &Handler{
		Deployer:       deployer,
		Compiler:       compiler,
		Verifier:       verifier,
		ArgsEncoder:    argsEncoder,
		CompilerJson:   compilerJson,
		InputGenerator: inputGenerator,
	}
}

// setup router
func (h *Handler) InitRouts() *gin.Engine {
	router := gin.Default()

	contractRoutes := router.Group("/contract")
	{
		contractRoutes.POST("/deploy", h.deploy)
		contractRoutes.POST("/abi", h.getABI)
		contractRoutes.POST("/bytecode", h.getBytecode)
		contractRoutes.POST("/verify", h.verify)
		contractRoutes.POST("/estimate-gas")
		contractRoutes.POST("/encode-function-call", h.encodeFunctionCall)
		contractRoutes.POST("/encode-constructor-args", h.encodeConstructorArgs)

	}

	//todo eth_estimateGas

	return router
}
