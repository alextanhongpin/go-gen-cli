package main

import (
	"fmt"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "inits a new gen.yaml file",
	Action: func(c *cli.Context) error {
		cfg := gen.Config{
			Templates: []*gen.Template{
				{
					Name:        "hello",
					Description: "hello template",
					Actions: []*gen.Action{
						gen.NewAction("hello"),
						gen.NewAction("hello_test"),
					},
				},
			},
		}

		cfgPath := c.String("file")
		if err := c.WriteConfigIfNotExists(cfgPath); err != nil {
			return err
		}
		NewSuccess(fmt.Sprintf("success: %s written", cfgPath))
		return nil
	},
}
