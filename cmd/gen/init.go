package main

import (
	"github.com/alextanhongpin/go-gen-cli/pkg/gen"

	"github.com/urfave/cli/v2"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "inits a new gen.yaml file",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "dry-run",
			Usage: "prints to stdout if true",
		},
	},
	Action: func(c *cli.Context) (err error) {
		cfgPath := (c.String("file"))
		dryRun := (c.Bool("dry-run"))
		return gen.Init(cfgPath, dryRun)
	},
}
