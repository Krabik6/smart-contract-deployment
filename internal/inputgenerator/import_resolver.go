package inputgenerator

import (
	"path/filepath"
	"regexp"
)

func (c *InputGenerator) findAndAddImports(content string, directory string, sources map[string]Source) error {
	// Use regex to find import statements
	re := regexp.MustCompile(`import "(.*?)";`)
	matches := re.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		// Resolve the path of the imported file relative to the directory of the current file
		relativePath := match[1]
		if _, exists := c.Sources[relativePath]; !exists {
			content, err := c.readSolidityFile(relativePath)
			if err != nil {
				return err
			}
			c.Sources[relativePath] = Source{Content: content}
			err = c.findAndAddImports(content, filepath.Dir(relativePath), sources)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
