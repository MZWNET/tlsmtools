package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// splitFileImpl is the actual implementation of the split functionality.
func splitFileImpl(filename string, outputDir string) error {
	origFilename := filepath.Join(outputDir, fmt.Sprintf("%s.orig", filepath.Base(filename)))
	altFilename := filepath.Join(outputDir, fmt.Sprintf("%s.alt", filepath.Base(filename)))

	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("error creating directory %s: %v", outputDir, err)
	}

	origFile, err := os.Create(origFilename)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", origFilename, err)
	}
	defer origFile.Close()

	altFile, err := os.Create(altFilename)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", altFilename, err)
	}
	defer altFile.Close()

	inputFile, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", filename, err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "orig: ") {
			origFile.WriteString(strings.TrimSuffix(strings.TrimPrefix(line, "orig: "), "%K%P") + "\n")
		} else if strings.HasPrefix(line, "alt(en): ") {
			altFile.WriteString(strings.TrimSuffix(strings.TrimPrefix(line, "alt(en): "), "%K%P") + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %v", filename, err)
	}

	return nil
}
