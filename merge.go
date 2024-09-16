package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// mergeFiles merges the orig, alt, and text files into a single tl file
// with the specified formatting, ensuring that each nth line from each file
// corresponds to the same block in the output file.
func mergeFiles(origFilename, altFilename, textFilename string) error {
	tlFilename := strings.TrimSuffix(origFilename, ".orig")
	tlFile, err := os.Create(tlFilename)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", tlFilename, err)
	}
	defer tlFile.Close()

	// Open the files
	origFile, err := os.Open(origFilename)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", origFilename, err)
	}
	defer origFile.Close()

	altFile, err := os.Open(altFilename)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", altFilename, err)
	}
	defer altFile.Close()

	textFile, err := os.Open(textFilename)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", textFilename, err)
	}
	defer textFile.Close()

	origScanner := bufio.NewScanner(origFile)
	altScanner := bufio.NewScanner(altFile)
	textScanner := bufio.NewScanner(textFile)

	for {
		origOk := origScanner.Scan()
		altOk := altScanner.Scan()
		textOk := textScanner.Scan()

		// Check if all files have been read completely
		if !origOk && !altOk && !textOk {
			break
		}

		if origOk {
			tlFile.WriteString(fmt.Sprintf("orig: %s%%K%%P\n", origScanner.Text()))
		}

		if altOk {
			tlFile.WriteString(fmt.Sprintf("alt(en): %s%%K%%P\n", altScanner.Text()))
		}

		if textOk {
			tlFile.WriteString(fmt.Sprintf("text: %s%%K%%P\n", textScanner.Text()))
		}

		// Write an empty line to separate blocks
		if origOk || altOk || textOk {
			tlFile.WriteString("\n")
		}
	}

	// Check for errors in scanners
	if err := origScanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %v", origFilename, err)
	}
	if err := altScanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %v", altFilename, err)
	}
	if err := textScanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %v", textFilename, err)
	}

	return nil
}
