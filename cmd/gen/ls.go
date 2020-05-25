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
		cfg, err := gen.NewConfig(c.String("file"))
		if err != nil {
			return err
		}

		for _, t := range cfg.Templates {
			fmt.Printf("%s: %s\n", t.Name, t.Description)
		}
		return nil
	},
}
