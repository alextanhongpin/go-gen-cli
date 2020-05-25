package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
)

var addCmd = &cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "adds a new template",
	Action: func(c *cli.Context) error {
		cfg, err := gen.NewConfig(c.String("file"))
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("error: %s is missing from the path", c.String("file"))
		}
		if err != nil {
			return err
		}

		name := c.Args().First()
		tpl := cfg.Find(name)
		if tpl == nil {
			tpl := gen.NewTemplate(name)
			tpl.Actions = append(tpl.Actions, gen.NewAction(name))
			cfg.Add(tpl)

			if err := gen.OverwriteConfig(cfg); err != nil {
				return err
			}
		}

		var errors []error
		for res := range gen.GenerateTemplates(tpl) {
			err, act := res.Error, res.Action
			if errors.Is(err, os.ErrExist) {
				NewWarning(fmt.Sprintf("error: file already exists at %s", act.Template))
				continue
			}
			if err != nil {
				errors = append(errors, err)
				continue
			}
			NewSuccess(fmt.Sprintf("success: created file %s", act.Template))
		}

		return cli.NewMultiError(errors...)
	},
}
