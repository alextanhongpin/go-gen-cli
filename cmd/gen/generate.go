package main

import (
	"bytes"
	"errors"
	"html/template"
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
		f, err := gen.Open(cfgPath, os.O_RDONLY)
		if err != nil {
			return err
		}
		defer f.Close()

		var cfg gen.Config
		if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
			return err
		}

		tplArg := c.Args().First()
		var tpl *gen.Template
		for _, t := range cfg.Templates {
			if t.Name == tplArg {
				tpl = t
				break
			}
		}
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
			// Format template and path name.
			var src, dst bytes.Buffer
			srctpl := template.Must(template.New("src").Parse(act.Template))
			_ = srctpl.Execute(&src, data)

			dsttpl := template.Must(template.New("dst").Parse(act.Path))
			_ = dsttpl.Execute(&dst, data)

			t, err := gen.Read(src.String())
			if err != nil {
				log.Printf("error reading template: %v\n", err)
				continue
			}
			if len(string(t)) == 0 {
				log.Printf("template is empty: %v\n", err)
				continue
			}
			if err := gen.Write(dst.String(), t, data); err != nil {
				log.Printf("error writing: %v\n", err)
				continue
			}
		}
		return nil
	},
}
