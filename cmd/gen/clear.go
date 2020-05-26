package main

import (
	"fmt"

	"github.com/alextanhongpin/go-gen-cli"

	"github.com/urfave/cli/v2"
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
			return fmt.Errorf("%s: not found", name)
		}

		merr := gen.NewMultiError()
		for _, a := range tpl.Actions {
			if !merr.Add(g.Remove(a.Template)) {
				fmt.Printf("%s: file removed\n", a.Template)
			}
		}
		return merr
	},
}
