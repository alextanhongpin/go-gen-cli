package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var generateCmd = &cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Usage:   "generates the given template",
	Action: func(c *cli.Context) error {
		b, err := gen.Read(cfgPath)
		if err != nil {
			return err
		}
		b = []byte(os.ExpandEnv(string(b)))

		var cfg gen.Config
		if err := yaml.Unmarshal(b, &cfg); err != nil {
			return err
		}

		name := c.Args().First()
		tpl := cfg.Find(name)
		if tpl == nil {
			return fmt.Errorf("error: template %q not found. templates available: %s", name, strings.Join(cfg.ListTemplates(), ", "))
		}

		// Prompt user for additional information.
		prompt, err := gen.Prompts(tpl.Prompts)
		if err != nil {
			return err
		}
		data.Prompt = prompt
		data.Env = tpl.Environment

		for key, value := range tpl.Environment {
			if gen.IsZero(value) {
				return fmt.Errorf("ENV %s is specified, but the value is empty", key)
			}
		}

		for _, act := range tpl.Actions {
			t, err := gen.Read(act.Template)
			if err != nil {
				NewWarning(fmt.Sprintf("error: %s", err))
				continue
			}
			if len(t) == 0 {
				NewWarning(fmt.Sprintf("error: template is empty, skipping %s", act.Template))
				continue
			}
			if err := gen.Write(act.Path, t, data); err != nil {
				if errors.Is(err, os.ErrExist) {
					NewWarning(fmt.Sprintf("error: file exists at %s", act.Path))
				} else {
					NewWarning(fmt.Sprintf("error: write failed: %s", err))
				}
				continue
			}
		}
		return nil
	},
}
