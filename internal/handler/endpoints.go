package handler

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Krabik6/smart-contract-deployment/pkg/api"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	gas, err := h.Deployer.EstimateGas(req.SourceCode, req.Optimize, req.Runs, req.ConstructorArguments)
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

	var args []interface{}
	if len(req.ConstructorArguments) > 0 {
		if err := json.Unmarshal(req.ConstructorArguments, &args); err != nil {
			c.JSON(400, gin.H{"error": "failed to parse ConstructorArguments"})
			return
		}
	}
	fmt.Printf("Arguments: %v\n", args)
	err := h.Verifier.Verify(req.ContractAddress, req.SourceCode, req.ContractName, req.LicenseType, req.Compilerversion, req.Optimize, req.Runs, args)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": "ok"})
}

// encodeConstructorArgs
func (h *Handler) encodeConstructorArgs(c *gin.Context) {
	var req api.EncodeConstructorArgsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var args []interface{}
	if len(req.Arguments) > 0 {
		if err := json.Unmarshal(req.Arguments, &args); err != nil {
			c.JSON(400, gin.H{"error": "failed to parse ConstructorArguments"})
			return
		}
	}

	encoded, err := h.Compiler.EncodeConstructorArgs(req.SourceCode, args...)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"encoded": hexutil.Encode(encoded)})
}

// encodeFunctionCall
func (h *Handler) encodeFunctionCall(c *gin.Context) {
	var req api.EncodeFunctionCallRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var args []interface{}
	if len(req.Arguments) > 0 {
		if err := json.Unmarshal(req.Arguments, &args); err != nil {
			c.JSON(400, gin.H{"error": "failed to parse ConstructorArguments"})
			return
		}
	}

	encoded, err := h.Compiler.EncodeFunctionCall(req.SourceCode, req.FunctionName, args...)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"encoded": hexutil.Encode(encoded)})
}
