package handler

import (
	"encoding/hex"
	"encoding/json"
	"github.com/Krabik6/smart-contract-deployment/internal/verify"
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
	//todo path
	input, path, err := h.InputGenerator.GenerateJSONInput("smart_contracts/smart.sol", req.Optimize, req.Runs)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// bytecode
	bytecode, err := h.CompilerJson.GetBytecode(input, path, req.ContractName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	//print bytecode in hex without 0x
	log.Println(hex.EncodeToString(bytecode))
	//abi
	abi, err := h.CompilerJson.GetAbi(input, path, req.ContractName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	contract, err := h.Deployer.Deploy(bytecode, abi, args...)
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

	if req.OptimizationUsed == nil {
		*req.OptimizationUsed = false
	}
	if req.Runs == nil {
		*req.Runs = 200
	}

	//todo path
	input, path, err := h.InputGenerator.GenerateJSONInput("smart_contracts/smart.sol", true, 200)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	//// bytecode
	//bytecode, err := h.CompilerJson.GetBytecode(input, path, req.ContractName)
	//if err != nil {
	//	c.JSON(500, gin.H{"error": err.Error()})
	//	return
	//}

	//abi
	abi, err := h.CompilerJson.GetAbi(input, path, "PublicStorageFuck")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	bytecode, err := h.CompilerJson.GetBytecode(input, path, "PublicStorageFuck")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	log.Println(hex.EncodeToString(bytecode))

	params := verify.Params{
		//APIKey:           "AFEMDPHAWXPHKI8SQJK9AS77UIAZN9NGCN",
		ContractAddress:  req.ContractAddress,
		SourceCode:       string(input),
		CodeFormat:       req.CodeFormat,
		ContractName:     req.ContractName,
		CompilerVersion:  req.CompilerVersion,
		OptimizationUsed: req.OptimizationUsed,
		Runs:             req.Runs,
		EVMVersion:       req.EVMVersion,
		LicenseType:      req.LicenseType,
		LibraryName1:     req.LibraryName1,
		LibraryAddress1:  req.LibraryAddress1,
		LibraryName2:     nil,
		LibraryAddress2:  nil,
		LibraryName3:     nil,
		LibraryAddress3:  nil,
		LibraryName4:     nil,
		LibraryAddress4:  nil,
		LibraryName5:     nil,
		LibraryAddress5:  nil,
		LibraryName6:     nil,
		LibraryAddress6:  nil,
		LibraryName7:     nil,
		LibraryAddress7:  nil,
		LibraryName8:     nil,
		LibraryAddress8:  nil,
		LibraryName9:     nil,
		LibraryAddress9:  nil,
		LibraryName10:    nil,
		LibraryAddress10: nil,
	}

	//params.SourceCode = inputStr
	params.CodeFormat = "solidity-standard-json-input"

	err = h.Verifier.Verify(abi, params, args...)
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

	abi, err := h.Compiler.GetAbi(req.SourceCode)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	encoded, err := h.ArgsEncoder.EncodeConstructorArgs(abi, args...)
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

	abi, err := h.Compiler.GetAbi(req.SourceCode)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	encoded, err := h.ArgsEncoder.EncodeConstructorArgs(abi, args...)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"encoded": hexutil.Encode(encoded)})
}
