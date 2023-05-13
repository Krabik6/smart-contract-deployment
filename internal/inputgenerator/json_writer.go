package inputgenerator

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type Source struct {
	Content string `json:"content"`
}

type Optimizer struct {
	Enabled bool `json:"enabled"`
	Runs    int  `json:"runs,omitempty"`
}

type Settings struct {
	OutputSelection map[string]map[string][]string `json:"outputSelection"`
	Optimizer       Optimizer                      `json:"optimizer,omitempty"`
}

type Compiler struct {
	Language string            `json:"language"`
	Sources  map[string]Source `json:"sources"`
	Settings Settings          `json:"settings"`
}

func NewCompiler() *Compiler {
	outputSelection := make(map[string]map[string][]string)
	outputSelection["*"] = make(map[string][]string)
	outputSelection["*"]["*"] = []string{"evm.bytecode", "evm.deployedBytecode", "abi"}

	return &Compiler{
		Language: "Solidity",
		Sources:  make(map[string]Source),
		Settings: Settings{OutputSelection: outputSelection},
	}
}

func (c *Compiler) WriteJSONInput(mainSolPath, outputPath string, optimize bool, optimizeRuns int) error {
	// Set the optimize option and optimization runs
	c.Settings.Optimizer.Enabled = optimize
	c.Settings.Optimizer.Runs = optimizeRuns

	// Read the main file
	mainContent, err := c.readSolidityFile(mainSolPath)
	if err != nil {
		return err
	}

	mainDirectory := filepath.Dir(mainSolPath)
	mainSolPath = strings.ReplaceAll(mainSolPath, "\\", "/") // Замена всех обратных слешей на прямые слеши
	c.Sources[mainSolPath] = Source{Content: mainContent}

	// Find and read imports
	err = c.findAndAddImports(mainContent, mainDirectory, c.Sources)
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(outputPath, jsonData, 0644)
	return err
}
