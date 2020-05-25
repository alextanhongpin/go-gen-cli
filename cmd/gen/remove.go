package main

import (
	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
)

var removeCmd = &cli.Command{
	Name:    "remove",
	Aliases: []string{"rm"},
	Usage:   "removes a registered template and all the generated files",
	Action: func(c *cli.Context) error {
		g := gen.New()
		err := g.Read(c.String("file"))
		if err != nil {
			return err
		}

		name := c.Args().First()
		tpl := g.FindTemplate(name)
		if tpl == nil {
			return nil
		}
		var errors Errors
		for _, act := range tpl.Actions {
			if err := act.RemoveTemplate(); err != nil {
				errors = append(errors, err)
			}
			if err := act.RemoveGeneratedFile(); err != nil {
				errors = append(errors, err)
			}
		}
		if len(errors) > 0 {
			return errors
		}

		g.RemoveTemplate(name)
		return g.Write(c.String("file"))
	},
}
