package main

import (
	"fmt"

	"github.com/alextanhongpin/go-gen"

	"github.com/urfave/cli"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "inits a new gen.yaml file",
	Action: func(c *cli.Context) (err error) {
		g := gen.New(c.String("file"))
		if err := g.Touch(c.String("file")); err != nil {
			return err
		}

		name := "hello"
		cfg := gen.NewConfig()
		cfg.Add(gen.NewTemplate(name))

		tpl := cfg.Find(name)

		merr := gen.NewMultiError()
		for _, a := range tpl.Actions {
			if !merr.Add(g.Touch(a.Template)) {
				fmt.Printf("%s: template created\n", a.Template)
			}
		}

		if !merr.Add(g.WriteConfig(cfg)) {
			fmt.Printf("%s: config written\n", c.String("file"))
		}

		return merr
	},
}
