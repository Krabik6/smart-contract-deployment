package compiler

import (
	"bufio"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type Compiler struct {
	WorkDir   string
	Image     string
	OutputDir string
}

func NewCompiler(workDir, image, outputDir string) *Compiler {
	return &Compiler{
		WorkDir:   workDir,
		Image:     image,
		OutputDir: outputDir,
	}
}

func (c *Compiler) Compile(file string) error {
	absPath, err := filepath.Abs(file)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for file %s: %v", file, err)
	}

	cmd := exec.Command("docker", "run", "-v", fmt.Sprintf("%s:/contract", filepath.Dir(absPath)), "ethereum/solc:stable", "--abi", "--bin", "-o", "/contract/"+c.OutputDir, "./contract/"+filepath.Base(absPath), "--overwrite")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to compile Solidity contract: %v", err)
	}
	return nil
}

func (c *Compiler) CompileSource(sourceCode string) error {
	if c.WorkDir == "" {
		return fmt.Errorf("work directory is not set")
	}
	if c.OutputDir == "" {
		return fmt.Errorf("output directory is not set")
	}
	fmt.Printf("Compiling Solidity contract in %s\n", c.WorkDir)
	cmd := exec.Command("docker", "run", "-i", "-v", fmt.Sprintf("%s:/contract", c.WorkDir), "ethereum/solc:stable", "--abi", "--bin", "-o", "/contract/"+c.OutputDir, "-", "--overwrite")
	cmd.Stdin = strings.NewReader(sourceCode)

	// выводим процесс компиляции в консоль
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start Solidity compiler: %v", err)
	}
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to compile Solidity contract: %v", err)
	}

	fmt.Println("Compiled successfully")
	return nil
}
