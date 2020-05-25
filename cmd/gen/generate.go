package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
)

var generateCmd = &cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Usage:   "generates the given template",
	Action: func(c *cli.Context) error {
		g := gen.New()
		err := g.Read(c.String("file"))
		if err != nil {
			return err
		}

		name := c.Args().First()
		tpl := g.FindTemplate(name)
		if tpl == nil {
			return fmt.Errorf("%s: not found", name)
		}

		// Prompt user for additional information.
		prompt, err := gen.Prompts(tpl.Prompts)
		if err != nil {
			return err
		}

		data := Data{
			Prompt: prompt,
			Env:    tpl.Environment,
		}

		errs := Errors(tpl.ValidateEnvironment())
		if len(errs) > 0 {
			return errs
		}

		errs = []error{}
		for _, act := range tpl.Actions {
			err := act.Exec(data)
			if errors.Is(err, os.ErrExist) {
				gen.Error("%s: file exists", act.Path)
				continue
			}
			if errors.Is(err, gen.ErrEmpty) {
				gen.Error("%s: file is empty", act.Template)
				continue
			}
			if err != nil {
				errs = append(errs, err)
				continue
			}
			gen.Info("%s: file created", act.Path)
		}
		return errs
	},
}
