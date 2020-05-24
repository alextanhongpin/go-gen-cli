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
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "struct",
			Usage:       "Sets the struct name",
			Destination: &data.StructName,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "type",
			Usage:       "Sets the struct type",
			Destination: &data.StructType,
			Required:    true,
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

		command := c.Args().First()
		var cmd *gen.Command
		for _, c := range cfg.Commands {
			if c.Name == command {
				cmd = c
				break
			}
		}
		if cmd == nil {
			return errors.New("command not found")
		}
		log.Println(cmd.Template)

		t, err := gen.Read(cmd.Template)
		if err != nil {
			return err
		}
		log.Println(string(t))
		if len(string(t)) == 0 {
			return errors.New("template is empty")
		}

		return gen.Write(cmd.Path, t, data)
	},
}
