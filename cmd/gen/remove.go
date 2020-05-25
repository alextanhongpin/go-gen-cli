package main

import (
	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var removeCmd = &cli.Command{
	Name:    "remove",
	Aliases: []string{"rm"},
	Usage:   "removes a registered template and all the generated files",
	Action: func(c *cli.Context) error {
		cfg, err := gen.NewConfig(c.String("file"))
		if err != nil {
			return err
		}

		name := c.Args().First()
		tpl := cfg.Find(name)
		for _, act := range tpl.Actions {
			// Format template and path name.
			if err := gen.RemoveIfExists(act.Template); err != nil {
				return err
			}
			if err := gen.RemoveIfExists(act.Path); err != nil {
				return err
			}
		}

		_ = cfg.Remove(name)

		b, err = yaml.Marshal(&cfg)
		if err != nil {
			return err
		}
		if err := gen.Overwrite(cfgPath, b); err != nil {
			return err
		}

		return nil
	},
}
