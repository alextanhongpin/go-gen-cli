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

var data Data

type Data struct {
	PackageName string
	StructName  string
	StructType  string
}

var templateCmd = &cli.Command{
	Name:    "template",
	Aliases: []string{"t"},
	Usage:   "options for templates",
	Subcommands: []*cli.Command{
		{
			Name:  "add",
			Usage: "add a new template",
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

				command := c.Args().First()
				if len(command) == 0 {
					return errors.New("template name cannot be empty")
				}
				var cmd *gen.Command
				for _, c := range cfg.Commands {
					if c.Name == command {
						cmd = c
						break
					}
				}
				log.Println("adding template", command)
				if cmd == nil {
					cmd = &gen.Command{
						Name:        command,
						Description: fmt.Sprintf("%s template", command),
						Template:    fmt.Sprintf("templates/%s.go", command),
						Path:        fmt.Sprintf("pkg/%s.go", command),
					}
					cfg.Commands = append(cfg.Commands, cmd)
					log.Println("adding config controller", cfg)

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

				log.Println(cmd)
				t, err := gen.Read(cmd.Template)
				if err != nil {
					return err
				}
				log.Println(t)
				if len(string(t)) == 0 {
					return errors.New("template is empty")
				}
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "remove a template",
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

				command := c.Args().First()
				if len(command) == 0 {
					return errors.New("template name cannot be empty")
				}
				var cmd *gen.Command
				var i int
				for i, cmd = range cfg.Commands {
					if cmd.Name == command {
						break
					}
				}
				if cmd == nil {
					return errors.New("no command found")
				}

				cfg.Commands = append(cfg.Commands[:i], cfg.Commands[i+1:]...)

				if err := f.Truncate(0); err != nil {
					return err
				}
				if _, err := f.Seek(0, 0); err != nil {
					return err
				}
				if err := yaml.NewEncoder(f).Encode(&cfg); err != nil {
					return err
				}
				if err := os.Remove(cmd.Template); err != nil {
					return err
				}
				if err := os.Remove(cmd.Path); err != nil {
					return err
				}
				log.Println("remove", cmd.Template)
				log.Println("remove", cmd.Path)
				return nil
			},
		},
	},
}
