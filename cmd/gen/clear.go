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
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "template",
			Aliases:  []string{"t"},
			Usage:    "use the given template",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		g := gen.New(c.String("file"))
		cfg, err := g.LoadConfig()
		if err != nil {
			return err
		}

		name := c.String("template")
		tpl := cfg.Find(name)
		if tpl == nil {
			return fmt.Errorf("%s: not found", name)
		}
		data := make(map[string]string, 0)
		data["Pkg"] = c.Args().First()

		merr := gen.NewMultiError()
		for _, vol := range tpl.Volumes {
			_, dst, err := vol.Split()
			if merr.Add(err) {
				continue
			}

			dst, err = gen.ParseString(dst, data)

			if merr.Add(err) {
				continue
			}
			if !merr.Add(g.Remove(dst)) {
				fmt.Printf("%s: file removed\n", dst)
			}
		}
		if merr.HasError() {
			return merr
		}
		return nil
	},
}
