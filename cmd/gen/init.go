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
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "multiple", Aliases: []string{"m"}},
	},
	Action: func(c *cli.Context) error {
		// Write-only, and create only if it does not exist.
		w, err := gen.Open(cfgPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL)
		if err != nil {
			return err
		}
		defer w.Close()

		var cfg gen.Config

		var tpl gen.Template
		if c.Bool("multiple") {
			tpl = gen.Template{
				Name:        "test",
				Description: "test template",
				Actions: []*gen.Action{
					{
						Description: "generate test",
						Template:    "templates/test.go",
						Path:        "pkg/test.go",
					},
					{
						Description: "generate test",
						Template:    "templates/test_test.go",
						Path:        "pkg/test.go",
					},
				},
			}
		} else {
			tpl = gen.Template{
				Name:        "test",
				Description: "test template",
				Template:    "templates/test.go",
				Path:        "pkg/test.go",
			}
		}
		cfg.Templates = append(cfg.Templates, &tpl)

		d, err := yaml.Marshal(&cfg)
		if err != nil {
			return err
		}

		n, err := w.Write(d)
		if err != nil {
			return err
		}
		fmt.Printf("Wrote config to %s (%d bytes)", cfgPath, n)
		return nil
	},
}
