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
	Deployer Deployer
	Compiler Compiler
}

func NewHandler(deployer Deployer, compiler Compiler) *Handler {
	return &Handler{
		Deployer: deployer,
		Compiler: compiler,
	}
}

// setup router
func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/deploy", h.deploy)
	router.GET("/abi", h.getABI)
	router.GET("/bytecode", h.getBytecode)
	router.GET("/abiBytecode", h.getABIAndBytecode)
	return router
}
