package main

import (
	"fmt"
	"os"

	"github.com/alextanhongpin/go-gen-cli"

	"github.com/urfave/cli/v2"
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
		tpl := cfg.Find(name)
		if tpl == nil {
			tpl = gen.NewTemplate(name)
			cfg.Add(tpl)

			if err := g.WriteConfig(cfg); err != nil {
				return err
			}
			fmt.Printf("%s: template added\n", name)
		}

		merr := gen.NewMultiError()
		for _, a := range tpl.Actions {
			path := os.ExpandEnv(a.Template)
			if !merr.Add(g.Touch(path)) {
				fmt.Printf("%s: file created\n", path)
			}
		}

		return merr
	},
}
