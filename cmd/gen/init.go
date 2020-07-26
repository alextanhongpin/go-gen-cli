package main

import (
	"fmt"

	"github.com/alextanhongpin/go-gen-cli"

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
		g := gen.New(c.String("file"))
		g.SetDryRun(c.Bool("dry-run"))

		name := "hello"
		cfg := gen.NewConfig()

		merr := gen.NewMultiError()

		if err := g.Touch(c.String("file")); err != nil {
			return err
		}

		tpl := cfg.Find(name)
		for _, vol := range tpl.Volumes {
			src, _, err := vol.Split()
			if merr.Add(err) {
				continue
			}

			if !merr.Add(g.Touch(src)) {
				fmt.Printf("%s: template created\n", src)
			}
		}

		if !merr.Add(g.WriteConfig(cfg)) {
			fmt.Printf("%s: config written\n", c.String("file"))
		}

		return merr
	},
}
