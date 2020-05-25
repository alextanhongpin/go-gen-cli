package main

import (
	"fmt"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "inits a new gen.yaml file",
	Action: func(c *cli.Context) error {
		cfg := gen.Config{
			Templates: []*gen.Template{
				{
					Name:        "hello ",
					Description: "hello template",
					Actions: []*gen.Action{
						gen.NewAction("hello"),
						gen.NewAction("hello_test"),
					},
				},
			},
		}

		d, err := yaml.Marshal(&cfg)
		if err != nil {
			return err
		}
		if err := gen.Write(cfgPath, d, nil); err != nil {
			return err
		}
		fmt.Printf("Wrote config to %s", cfgPath)
		return nil
	},
}
