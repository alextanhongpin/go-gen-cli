package main

import (
	"fmt"

	"github.com/alextanhongpin/go-gen"

	"github.com/urfave/cli"
)

var addCmd = &cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "adds a new template",
	Action: func(c *cli.Context) error {
		g := gen.New(c.String("file"))
		cfg, err := g.LoadConfig()
		if err != nil {
			return err
		}

		name := c.Args().First()
		fmt.Println("addingtemplate")
		tpl := cfg.Find(name)
		if tpl == nil {
			fmt.Println("template not found, creating")
			tpl = gen.NewTemplate(name)
			cfg.Add(tpl)

			if err := g.WriteConfig(cfg); err != nil {
				return err
			}
		}

		merr := gen.NewMultiError()
		for _, a := range tpl.Actions {
			merr.Add(g.Touch(a.Template))
		}
		return merr
	},
}
