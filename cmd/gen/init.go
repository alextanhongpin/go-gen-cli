package main

import (
	"fmt"

	"github.com/alextanhongpin/go-gen-cli"

	"github.com/urfave/cli"
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
		g := gen.New(c.String("file"))
		g.SetDryRun(c.Bool("dry-run"))

		name := "hello"
		cfg := gen.NewConfig()
		cfg.Add(gen.NewTemplate(name))

		merr := gen.NewMultiError()

		if err := g.Touch(c.String("file")); err != nil {
			return err
		}

		tpl := cfg.Find(name)
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
