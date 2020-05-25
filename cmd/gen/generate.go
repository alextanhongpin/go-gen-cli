package main

import (
	"errors"
	"log"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var generateCmd = &cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Usage:   "generates the given template",
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
		rawCfg := gen.Read(cfgPath)
		rawCfg = []byte(os.ExpandEnv(string(rawCfg)))

		var cfg gen.Config
		if err := yaml.Unmarshal(rawCfg, &cfg); err != nil {
			return err
		}

		name := c.Args().First()
		tpl := cfg.Find(name)
		if tpl == nil {
			return errors.New("template not found")
		}
		data.Template = tpl

		// Prompt user for additional information.
		prompts, err := gen.Prompts(tpl.Prompts)
		if err != nil {
			return err
		}
		data.Prompts = prompts

		for _, act := range tpl.Actions {
			t, err := gen.Read(act.Template)
			if err != nil {
				log.Printf("error reading template: %v\n", err)
				continue
			}
			if len(string(t)) == 0 {
				log.Printf("template is empty: %v\n", err)
				continue
			}
			if err := gen.Write(act.Path, t, data); err != nil {
				log.Printf("error writing: %v\n", err)
				continue
			}
		}
		return nil
	},
}
