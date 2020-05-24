package main

import (
	"fmt"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var lsCmd = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "lists the existing templates",
	Action: func(c *cli.Context) error {
		f, err := gen.Open(cfgPath, os.O_RDONLY)
		if err != nil {
			return err
		}
		defer f.Close()

		var cfg gen.Config
		if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
			return err
		}

		for _, tpl := range cfg.Templates {
			fmt.Printf("- name: %s\n", tpl.Name)
			fmt.Printf("  description: %s\n", tpl.Description)
		}
		return nil
	},
}
