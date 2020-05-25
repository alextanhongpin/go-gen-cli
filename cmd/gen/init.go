package main

import (
	"fmt"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
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
					// Actions: []*gen.Action{
					//         gen.NewAction("hello"),
					//         gen.NewAction("hello_test"),
					// },
				},
			},
		}

		b, err := yaml.Marshal(&cfg)
		if err != nil {
			return err
		}
		cfgPath := c.String("file")
		f, err := gen.Open(cfgPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL)
		if err != nil {
			if errors.Is(err, os.ErrExist) {
				return fmt.Errorf("error: %s already exists", cfgPath)
			}
			return err
		}
		defer f.Close()

		if _, err := f.Write(b); err != nil {
			return err
		}
		NewSuccess(fmt.Sprintf("success: %s written", cfgPath))
		return nil
	},
}
