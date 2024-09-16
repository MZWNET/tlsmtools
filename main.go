package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "tlsmtools",
		Usage: "Split and merge tl files",
		Commands: []*cli.Command{
			{
				Name:    "split",
				Aliases: []string{"s"},
				Usage:   "Split the given file into orig and alt parts",
				Action:  splitFile,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "File to split",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "outputDir",
						Aliases: []string{"d"},
						Usage:   "Directory to save the split files",
						Value:   ".",
					},
				},
			},
			{
				Name:    "merge",
				Aliases: []string{"m"},
				Usage:   "Merge orig, alt, and text files into a single tl file",
				Action:  mergeAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "orig",
						Usage:    "orig file",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "alt",
						Usage:    "alt file",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "text",
						Usage:    "text file",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "Output file name and path",
						Value:   "output.tl",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func splitFile(c *cli.Context) error {
	filename := c.String("file")
	outputDir := c.String("outputDir")
	return splitFileImpl(filename, outputDir)
}

func mergeAction(c *cli.Context) error {
	origFile := c.String("orig")
	altFile := c.String("alt")
	textFile := c.String("text")
	outputFile := c.String("output")
	return mergeFiles(origFile, altFile, textFile, outputFile)
}
