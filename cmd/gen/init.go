package main

import (
	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "inits a new gen.yaml file",
	Action: func(c *cli.Context) (err error) {
		defer func() {
			if err == nil {
				gen.Info("%s: config written", c.String("file"))
			}
		}()
		if err = gen.Touch(c.String("file")); err != nil {
			return err
		}

		g := gen.New()
		g.AddTemplate(gen.NewTemplate("hello_world"))
		return g.Write(c.String("file"))
	},
}
