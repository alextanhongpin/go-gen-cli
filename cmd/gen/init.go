package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "inits a new gen.yaml file",
	Action: func(c *cli.Context) error {
		g := gen.New()
		g.AddTemplate(gen.NewTemplate("hello_world"))
		err := g.Write(c.String("file"))
		if errors.Is(err, os.ErrExist) {
			return errors.New(fmt.Sprintf("error: config exists %s", c.String("file")))
		}
		if err != nil {
			return err
		}

		NewSuccess(fmt.Sprintf("success: config generated at %s", c.String("file")))
		return nil
	},
}
