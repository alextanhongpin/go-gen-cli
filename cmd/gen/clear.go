package main

import (
	"errors"

	"github.com/alextanhongpin/go-gen"

	"github.com/urfave/cli"
)

var clearCmd = &cli.Command{
	Name:    "clear",
	Aliases: []string{"c"},
	Usage:   "clears the generated files for a given template",
	Action: func(c *cli.Context) error {
		g := gen.New(c.String("file"))
		cfg, err := g.LoadConfig()
		if err != nil {
			return err
		}

		name := c.Args().First()
		tpl := cfg.Find(name)
		if tpl == nil {
			return errors.New("template: not found")
		}

		merr := gen.NewMultiError()
		for _, a := range tpl.Actions {
			merr.Add(g.Remove(a.Template))
		}
		return merr
	},
}
