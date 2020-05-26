package main

import (
	"github.com/alextanhongpin/go-gen"

	"github.com/urfave/cli"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "inits a new gen.yaml file",
	Action: func(c *cli.Context) (err error) {
		g := gen.New(c.String("file"))
		g.Touch(c.String("file"))

		cfg := gen.NewConfig()
		cfg.Add(gen.NewTemplate("hello"))
		return g.WriteConfig(cfg)
	},
}
