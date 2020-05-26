package main

import (
	"github.com/alextanhongpin/go-gen"

	"github.com/urfave/cli"
)

var removeCmd = &cli.Command{
	Name:    "remove",
	Aliases: []string{"rm"},
	Usage:   "removes a registered template and all the generated files",
	Action: func(c *cli.Context) error {
		g := gen.New(c.String("file"))
		cfg, err := g.LoadConfig()
		if err != nil {
			return err
		}

		name := c.Args().First()
		tpl := cfg.Find(name)
		if tpl == nil {
			return nil
		}

		merr := gen.NewMultiError()
		for _, act := range tpl.Actions {
			merr.Add(g.Remove(act.Template))
			merr.Add(g.Remove(act.Path))
		}

		cfg.Remove(name)

		merr.Add(g.WriteConfig(cfg))

		return merr
	},
}
