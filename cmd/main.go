package main

import (
	"os"

	"github.com/urfave/cli/v3"
)

var app = new(cli.App)
var fileInput string

func main() {
	app.Commands = []*cli.Command{
		{
			Name:        "init",
			Usage:       "generate example collections.yaml, if exsisted, will do nothing",
			Description: "generate example collections.yaml, if exsisted, will do nothing",
			Action:      initCmd,
		},
		{
			Name:        "generate",
			Usage:       "generate go file from input.yaml, default is collections.yaml",
			Description: "generate go file from input.yaml, default is collections.yaml",
			Action:      generate,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "file",
					Usage:       "input file path, default is collections.yaml",
					Destination: &fileInput,
					Value:       "collections.yaml",
					Aliases:     []string{"f"},
				},
			},
		},
		{
			Name:   "version",
			Usage:  "Show  app version",
			Action: version,
		},
	}

	app.Run(os.Args)
}
