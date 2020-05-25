package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

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
			return fmt.Errorf("error: template %q not found. templates available: %s", name, strings.Join(g.ListTemplates(), ", "))
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
				NewWarning(fmt.Sprintf("error: file exists at %s", act.Path))
				continue
			}
			if errors.Is(err, gen.ErrEmpty) {
				NewWarning(fmt.Sprintf("error: file is empty %s", act.Template))
				continue
			}
			if err != nil {
				errs = append(errs, err)
				continue
			}
			NewSuccess(fmt.Sprintf("success: wrote to %s", act.Path))
		}
		return errs
	},
}
