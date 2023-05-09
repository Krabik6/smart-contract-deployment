package handler

import (
	"encoding/hex"
	"encoding/json"
	"github.com/Krabik6/smart-contract-deployment/internal/verify"
	"github.com/Krabik6/smart-contract-deployment/pkg/api"
	"github.com/gin-gonic/gin"
	"log"
)

// deploy handler
func (h *Handler) deploy(c *gin.Context) {
	var req api.DeployRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var args []interface{}
	if len(req.ConstructorArguments) > 0 {
		if err := json.Unmarshal(req.ConstructorArguments, &args); err != nil {
			c.JSON(400, gin.H{"error": "failed to parse ConstructorArguments"})
			return
		}
	}

	contract, err := h.Deployer.Deploy(req.SourceCode, req.Optimize, req.Runs, args...)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"address": contract})
}

// getABI handler
func (h *Handler) getABI(c *gin.Context) {
	var req api.AbiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println(req)
	abi, err := h.Compiler.GetAbi(req.SourceCode)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"abi": abi})
}

// getBytecode handler
func (h *Handler) getBytecode(c *gin.Context) {
	var req api.BytecodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println(req.SourceCode)

	bytecode, err := h.Compiler.GetBytecode(req.SourceCode, req.Optimize, req.Runs)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	//bytecode []byte to string
	bytecodeStr := hex.EncodeToString(bytecode)
	c.JSON(200, gin.H{"bytecode": bytecodeStr})
}

// EstimateGas handler
func (h *Handler) estimateGas(c *gin.Context) {
	var req api.DeployRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gas, err := h.Deployer.EstimateGas(req.SourceCode, req.ConstructorArguments)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"gas": gas})
}

// verify
func (h *Handler) verify(c *gin.Context) {
	var req api.VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := verify.Verify(req.SourceCode, req.ContractAddress, req.ContractName, string(req.ConstructorArguments))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": "ok"})
}
