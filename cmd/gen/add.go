package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var addCmd = &cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "adds a new template",
	Action: func(c *cli.Context) error {
		b, err := gen.Read(cfgPath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("error: %s is missing from the path", cfgPath)
			}
			return err
		}
		b = []byte(os.ExpandEnv(string(b)))

		// Load as YAML.
		var cfg gen.Config
		if err := yaml.Unmarshal(b, &cfg); err != nil {
			return err
		}

		name := c.Args().First()
		tpl := cfg.Find(name)
		if tpl == nil {
			tpl = &gen.Template{
				Name:        name,
				Description: fmt.Sprintf("%s template", name),
			}
			tpl.Actions = []*gen.Action{gen.NewAction(name)}
			cfg.Add(tpl)

			b, err := yaml.Marshal(&cfg)
			if err != nil {
				return err
			}
			if err := gen.Overwrite(cfgPath, b); err != nil {
				return err
			}
		}

		for _, act := range tpl.Actions {
			if err := gen.Create(act.Template); err != nil {
				if errors.Is(err, os.ErrExist) {
					NewWarning(fmt.Sprintf("error: file already exists at %s", act.Template))
					continue
				} else {
					return err
				}
			}
			NewSuccess(fmt.Sprintf("success: created file %s", act.Template))
		}
		return nil
	},
}
