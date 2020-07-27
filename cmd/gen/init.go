package main

import (
	"fmt"

	"github.com/alextanhongpin/go-gen-cli"

	"github.com/urfave/cli/v2"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "inits a new gen.yaml file",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "dry-run",
			Usage: "prints to stdout if true",
		},
	},
	Action: func(c *cli.Context) (err error) {
		g := gen.New(c.String("file"))
		g.SetDryRun(c.Bool("dry-run"))

		cfg := gen.NewConfig()

		merr := gen.NewMultiError()

		if err := g.Touch(c.String("file")); err != nil {
			return err
		}

		cfg.Version = "0.1"
		cfg.Templates = append(cfg.Templates, &gen.Template{
			Name:    "domain",
			Volumes: []gen.Volume{"templates/controller.tmpl:/tmp/{{.Pkg}}/controller.go"},
			Environment: map[string]string{
				"Controller":        `{{ pascalcase .Pkg }}Controller`,
				"Entity":            `{{ pascalcase .Pkg }}`,
				"Service":           `{{ pascalcase "$PKG" }}Service`,
				"CamelCaseSingular": `{{ camelcase .Pkg }}`,
				"CamelCasePlural":   `{{ camelcase .Pkg }}s`,
				"SnakeCaseSingular": `{{ snakecase .Pkg }}`,
				"SnakeCasePlural":   `{{ snakecase .Pkg }}s`,
			},
		})

		tpl := cfg.Find("domain")
		for _, vol := range tpl.Volumes {
			src, _, err := vol.Split()
			if merr.Add(err) {
				continue
			}

			if !merr.Add(g.Touch(src)) {
				fmt.Printf("%s: template created\n", src)
			}
		}

		if !merr.Add(g.WriteConfig(cfg)) {
			fmt.Printf("%s: config written\n", c.String("file"))
		}

		if merr.HasError() {
			return merr
		}
		return nil
	},
}
