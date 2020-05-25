package main

import (
	"fmt"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var lsCmd = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "lists the existing templates",
	Action: func(c *cli.Context) error {
		b, err := gen.Read(c.String("file"))
		if err != nil {
			return err
		}

		var cfg gen.Config
		if err := yaml.Unmarshal(b, &cfg); err != nil {
			return err
		}

		for _, tpl := range cfg.Templates {
			fmt.Printf("- name: %s\n", tpl.Name)
			fmt.Printf("  description: %s\n", tpl.Description)
		}
		return nil
	},
}
