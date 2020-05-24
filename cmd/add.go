package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var addCmd = &cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "adds a new template",
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
		for _, t := range cfg.Templates {
			if t.Name == tplArg {
				tpl = t
				break
			}
		}
		if tpl == nil {
			tpl = &gen.Template{
				Name:        tplArg,
				Description: fmt.Sprintf("%s template", tplArg),
				Template:    fmt.Sprintf("templates/%s.go", tplArg),
				Path:        fmt.Sprintf("pkg/%s.go", tplArg),
			}
			cfg.Templates = append(cfg.Templates, tpl)

			if err := f.Truncate(0); err != nil {
				return err
			}
			if _, err := f.Seek(0, 0); err != nil {
				return err
			}
			if err := yaml.NewEncoder(f).Encode(&cfg); err != nil {
				return err
			}
		}

		t, err := gen.Read(tpl.Template)
		if err != nil {
			return err
		}
		if len(string(t)) == 0 {
			return errors.New("template is empty")
		}
		return nil
	},
}
