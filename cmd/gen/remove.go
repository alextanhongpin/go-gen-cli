package main

import (
	"bytes"
	"errors"
	"html/template"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/Masterminds/sprig"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var removeCmd = &cli.Command{
	Name:    "remove",
	Aliases: []string{"rm"},
	Usage:   "removes a template and the generated files and configuration",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "pkg",
			Usage:       "Sets the package name",
			Destination: &data.PackageName,
		},
		&cli.StringFlag{
			Name:        "struct",
			Usage:       "Sets the struct name",
			Destination: &data.StructName,
		},
		&cli.StringFlag{
			Name:        "tag",
			Usage:       "Sets a tag",
			Destination: &data.Tag,
		},
	},
	Action: func(c *cli.Context) error {
		f, err := gen.Open(cfgPath, os.O_RDWR)
		if err != nil {
			return err
		}
		defer f.Close()

		var cfg gen.Config
		if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
			return err
		}

		tplArg := c.Args().First()
		if len(tplArg) == 0 {
			return errors.New("template name cannot be empty")
		}

		var tpl *gen.Template
		var i int
		for j, t := range cfg.Templates {
			if t.Name == tplArg {
				i = j
				tpl = t
				break
			}
		}
		if tpl == nil {
			return errors.New("no template found")
		}

		cfg.Templates = append(cfg.Templates[:i], cfg.Templates[i+1:]...)

		if err := f.Truncate(0); err != nil {
			return err
		}
		if _, err := f.Seek(0, 0); err != nil {
			return err
		}
		if err := yaml.NewEncoder(f).Encode(&cfg); err != nil {
			return err
		}

		if len(tpl.Actions) > 0 {
			for _, act := range tpl.Actions {
				// Format template and path name.
				var src, dst bytes.Buffer
				srctpl := template.Must(template.New("src").Funcs(sprig.FuncMap()).Parse(act.Template))
				_ = srctpl.Execute(&src, data)

				dsttpl := template.Must(template.New("dst").Funcs(sprig.FuncMap()).Parse(act.Path))
				_ = dsttpl.Execute(&dst, data)

				if err := os.Remove(src.String()); !errors.Is(err, os.ErrNotExist) {
					return err
				}
				if err := os.Remove(dst.String()); !errors.Is(err, os.ErrNotExist) {
					return err
				}
			}
		}
		if err := os.Remove(tpl.Template); !errors.Is(err, os.ErrNotExist) {
			return err
		}
		if err := os.Remove(tpl.Path); !errors.Is(err, os.ErrNotExist) {
			return err
		}
		return nil
	},
}
