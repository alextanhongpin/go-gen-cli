package main

import (
	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
)

var clearCmd = &cli.Command{
	Name:    "clear",
	Aliases: []string{"c"},
	Usage:   "clears the generated files for a given template",
	Action: func(c *cli.Context) error {
		g := gen.New()
		if err := g.Read(c.String("file")); err != nil {
			return err
		}

		name := c.Args().First()
		tpl := g.FindTemplate(name)

		var errors Errors
		for _, act := range tpl.Actions {
			if err := act.RemoveGeneratedFile(); err != nil {
				errors = append(errors, err)
				continue
			}
			gen.Info("%s: file removed", act.Path)
		}

		return errors
	},
}
