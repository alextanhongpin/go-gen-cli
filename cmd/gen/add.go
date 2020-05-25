package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
)

type Errors []error

func (e Errors) Error() string {
	return ""
}

func (e Errors) Errors() []error {
	return []error(e)
}

var addCmd = &cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "adds a new template",
	Action: func(c *cli.Context) error {
		g := gen.New()
		err := g.Read(c.String("file"))
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("error: %s is missing from the path", c.String("file"))
		}
		if err != nil {
			return err
		}

		name := c.Args().First()
		tpl := g.FindTemplate(name)
		if tpl == nil {
			tpl = gen.NewTemplate(name)
			g.AddTemplate(tpl)

			if err := g.Write(c.String("file")); err != nil {
				return err
			}
		}

		var errs Errors
		for _, act := range tpl.Actions {
			err := act.TouchTemplate()
			if errors.Is(err, os.ErrExist) {
				NewWarning(fmt.Sprintf("error: file already exists at %s, skipping create", act.Template))
				continue
			}
			if err != nil {
				errs = append(errs, err)
				continue
			}
			NewSuccess(fmt.Sprintf("success: created file %s", act.Template))
		}

		return errs
	},
}
