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
						Aliases:  []string{"o"},
						Usage:    "Original file",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "alt",
						Aliases:  []string{"a"},
						Usage:    "Alternate file",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "text",
						Aliases:  []string{"t"},
						Usage:    "Text file",
						Required: true,
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
	return splitFileImpl(filename)
}

func mergeAction(c *cli.Context) error {
	origFile := c.String("orig")
	altFile := c.String("alt")
	textFile := c.String("text")
	return mergeFiles(origFile, altFile, textFile)
}
