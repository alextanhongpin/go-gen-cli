package main

import (
	"fmt"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var clearCmd = &cli.Command{
	Name:    "clear",
	Aliases: []string{"c"},
	Usage:   "clears the generated files for a given template",
	Action: func(c *cli.Context) error {
		b, err := gen.Read(cfgPath)
		if err != nil {
			return err
		}
		b = []byte(os.ExpandEnv(string(b)))

		var cfg gen.Config
		if err := yaml.Unmarshal(b, &cfg); err != nil {
			return err
		}

		name := c.Args().First()
		tpl := cfg.Find(name)

		for _, act := range tpl.Actions {
			if err := gen.RemoveIfExists(act.Path); err != nil {
				return err
			}
			NewInfo(fmt.Sprintf("info: remove generated template %s", act.Path))
		}

		return nil
	},
}
