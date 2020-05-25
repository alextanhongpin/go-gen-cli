package main

import (
	"errors"
	"fmt"
	"log"
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
			}
			tpl.Actions = []*gen.Action{
				{
					Description: fmt.Sprintf("%s template", tplArg),
					Template:    fmt.Sprintf("templates/%s.go", tplArg),
					Path:        fmt.Sprintf("pkg/%s.go", tplArg),
				},
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

		for _, act := range tpl.Actions {
			r, err := gen.Open(act.Template, os.O_RDONLY|os.O_CREATE|os.O_EXCL)
			defer r.Close()
			if errors.Is(err, os.ErrExist) {
				log.Printf("file exists: %s\n", act.Template)
				continue
			}
			if err != nil {
				return err
			}
			log.Printf("created file %s\n", act.Template)
		}
		return nil
	},
}
