package inputgenerator

import (
	"os"
)

func (c *Compiler) readSolidityFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
