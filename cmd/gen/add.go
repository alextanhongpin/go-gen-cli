package main

import (
	"errors"
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
				gen.Error("%s: file exists", act.Template)
				continue
			}
			if err != nil {
				errs = append(errs, err)
				continue
			}
			gen.Info("%s: file created", act.Template)
		}

		return errs
	},
}
