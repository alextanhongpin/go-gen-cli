package main

import (
	"errors"
	"fmt"

	"github.com/alextanhongpin/go-gen-cli"

	"github.com/urfave/cli/v2"
)

var clearCmd = &cli.Command{
	Name:    "clear",
	Aliases: []string{"c"},
	Usage:   "clears the generated files for a given template",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "template",
			Aliases:  []string{"t"},
			Usage:    "use the given template",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		configPath := c.String("file")
		templateName := c.String("template")
		packageName := c.Args().First()
		return Clear(configPath, templateName, packageName)
	},
}

func Clear(configPath, templateName, packageName string) error {
	g := gen.New(configPath)
	cfg, err := g.LoadConfig()
	if err != nil {
		return err
	}
	tpl := cfg.Find(templateName)
	if tpl == nil {
		return fmt.Errorf("%s: not found", templateName)
	}
	data := make(map[string]interface{}, 0)
	data["Pkg"] = packageName

	merr := gen.NewMultiError()
	for _, act := range tpl.Actions {
		newAct, err := act.Resolve(data)
		if merr.Add(err) {
			continue
		}

		dst := newAct.Destination()
		if dst == "" && merr.Add(errors.New("destination is empty")) {
			continue
		}

		if !merr.Add(g.Remove(dst)) {
			fmt.Printf("%s: file removed\n", dst)
		}
	}
	if merr.HasError() {
		return merr
	}
	return nil
}
