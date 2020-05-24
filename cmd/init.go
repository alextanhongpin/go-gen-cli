package main

import (
	"fmt"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "inits a new gen.yaml file",
	Action: func(c *cli.Context) error {
		// Write-only, and create only if it does not exist.
		w, err := gen.Open(cfgPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL)
		if err != nil {
			return err
		}
		defer w.Close()

		var cfg gen.Config
		cfg.Commands = append(cfg.Commands, &gen.Command{
			Name:        "test",
			Description: "test template",
			Template:    "templates/test.go",
			Path:        "pkg/test.go",
		})

		d, err := yaml.Marshal(&cfg)
		if err != nil {
			return err
		}

		n, err := w.Write(d)
		if err != nil {
			return err
		}
		fmt.Printf("Write %d bytes", n)
		return nil
	},
}
