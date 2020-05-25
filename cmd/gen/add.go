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
		// Open config file.
		f, err := gen.Open(cfgPath, os.O_RDWR)
		if err != nil {
			return err
		}
		defer f.Close()

		// Load as YAML.
		var cfg gen.Config
		if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
			return err
		}

		name := c.Args().First()
		tpl := cfg.Templates.Find(name)
		if tpl == nil {
			tpl = &gen.Template{
				Name:        name,
				Description: fmt.Sprintf("%s template", name),
			}
			tpl.Actions = []*gen.Action{gen.NewAction(name)}
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
			if err := gen.Create(act.Template); err != nil {
				if errors.Is(err, os.ErrExist) {
					log.Printf("file exists: %s\n", act.Template)
					continue
				} else {
					return err
				}
			}
			log.Printf("created file %s\n", act.Template)
		}
		return nil
	},
}
