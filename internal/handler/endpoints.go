package handler

import (
	"github.com/Krabik6/smart-contract-deployment/pkg/api"
	"github.com/gin-gonic/gin"
)

// deploy handler
func (h *Handler) deploy(c *gin.Context) {
	var req api.DeployRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	contract, err := h.Deployer.Deploy(req.SourceCode, req.ConstructorArguments)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"address": contract})
}

// getABI handler
func (h *Handler) getABI(c *gin.Context) {
	var req api.SourceCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	abi, err := h.Compiler.GetABI(req.SourceCode)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"abi": abi})
}

// getBytecode handler
func (h *Handler) getBytecode(c *gin.Context) {
	var req api.SourceCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	bytecode, err := h.Compiler.GetBytecode(req.SourceCode)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"bytecode": bytecode})
}

// getABIAndBytecode handler
func (h *Handler) getABIAndBytecode(c *gin.Context) {
	var req api.SourceCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	abi, bytecode, err := h.Compiler.GetABIAndBytecode(req.SourceCode)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"abi": abi, "bytecode": bytecode})
}
