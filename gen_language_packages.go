//go:build ignore

package main

import (
	"fmt"
	"os"
)

const modelDir = "language-models"

func main() {
	dirs, err := os.ReadDir(modelDir)
	if err != nil {
		fmt.Printf("Failed to read language model directory: %v\n", err)
		os.Exit(1)
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		langCode := dir.Name()
		dirPath := fmt.Sprintf("%s/%s", modelDir, langCode)

		filePath := fmt.Sprintf("%s/embed.go", dirPath)

		content := fmt.Sprintf(`package %s
import (
	"embed"
	"github.com/asciimoo/lingua-go"
)
//go:embed *.zip
var model embed.FS
func init() {
	lingua.Register("%s", model)
}
`, langCode, langCode)

		err = os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Failed to generate language package %s. Error: %v\n", filePath, err)
			continue
		}

		fmt.Printf("Success: %s\n", filePath)
	}
}
