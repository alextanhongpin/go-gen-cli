package main

import (
	"fmt"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
)

var lsCmd = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "lists the existing templates",
	Action: func(c *cli.Context) error {
		g := gen.New()
		err := g.Read(c.String("file"))
		if err != nil {
			return err
		}

		for _, t := range g.ListTemplates() {
			fmt.Println(t)
		}
		return nil
	},
}
