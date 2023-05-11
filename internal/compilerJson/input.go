package compilerjson

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type Source struct {
	Content string `json:"content"`
}

type Settings struct {
	OutputSelection map[string]map[string][]string `json:"outputSelection"`
}

type Input struct {
	Language string            `json:"language"`
	Sources  map[string]Source `json:"sources"`
	Settings Settings          `json:"settings"`
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func readSolidityFile(path string) string {
	data, err := os.ReadFile(path)
	checkError(err)
	return string(data)
}

func findImports(content string, directory string) []string {
	// Use regex to find import statements
	re := regexp.MustCompile(`import "(.*?)";`)
	matches := re.FindAllStringSubmatch(content, -1)

	// Collect paths
	var paths []string
	for _, match := range matches {
		// Resolve the path of the imported file relative to the directory of the current file
		relativePath := filepath.Join(directory, match[1])
		paths = append(paths, relativePath)
	}

	return paths
}

func WriteJSONInput(mainSolPath, outputPath string) {
	sources := make(map[string]Source)

	// Read the main file
	mainContent := readSolidityFile(mainSolPath)
	mainDirectory := filepath.Dir(mainSolPath)
	sources[filepath.Base(mainSolPath)] = Source{Content: mainContent}

	// Find and read imports
	imports := findImports(mainContent, mainDirectory)
	for _, importPath := range imports {
		content := readSolidityFile(importPath)
		sources[filepath.Base(importPath)] = Source{Content: content}
	}

	outputSelection := make(map[string]map[string][]string)
	outputSelection["*"] = make(map[string][]string)
	outputSelection["*"]["*"] = []string{"evm.bytecode", "evm.deployedBytecode", "abi"}

	input := Input{
		Language: "Solidity",
		Sources:  sources,
		Settings: Settings{OutputSelection: outputSelection},
	}

	jsonData, err := json.MarshalIndent(input, "", "  ")
	checkError(err)

	err = ioutil.WriteFile(outputPath, jsonData, 0644)
	checkError(err)
}
