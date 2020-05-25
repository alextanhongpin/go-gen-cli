package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "inits a new gen.yaml file",
	Action: func(c *cli.Context) error {
		cfg := gen.Config{
			Templates: []*gen.Template{
				{
					Name:        "hello",
					Description: "hello template",
					Actions: []*gen.Action{
						gen.NewAction("hello"),
						gen.NewAction("hello_test"),
					},
				},
			},
		}

		b, err := yaml.Marshal(&cfg)
		if err != nil {
			return err
		}
		f, err := gen.Open(cfgPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL)
		if err != nil {
			if errors.Is(err, os.ErrExist) {
				return NewError("gen.yaml already exists")
			}
			return err
		}
		defer f.Close()
		if _, err := f.Write(b); err != nil {
			return err
		}
		fmt.Println(gen.Success(fmt.Sprintf("%s written", cfgPath)))
		return nil
	},
}
