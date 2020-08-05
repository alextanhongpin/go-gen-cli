package main

import (
	"errors"
	"fmt"
	"log"
	"path"

	"github.com/alextanhongpin/go-gen-cli"

	"github.com/urfave/cli/v2"
)

var generateCmd = &cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Usage:   "generates the given template",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "dry-run",
			Usage: "prints to stdout if true",
		},
		&cli.StringFlag{
			Name:     "template",
			Aliases:  []string{"t"},
			Usage:    "use the given template",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		cfgPath := c.String("file")
		dryRun := c.Bool("dry-run")
		templateName := c.String("template")
		pkgName := c.Args().First()

		return Generate(cfgPath, pkgName, templateName, dryRun)
	},
}

func Generate(cfgPath, pkgName, templateName string, dryRun bool) error {
	g := gen.New(cfgPath)
	g.SetDryRun(dryRun)

	cfg, err := g.LoadConfig()
	if err != nil {
		return err
	}

	if templateName == "" {
		return errors.New("name is required: gen generate -t <template-name> <pkg-name>")
	}
	tpl := cfg.Find(templateName)
	if tpl == nil {
		return fmt.Errorf("%s: not found", templateName)
	}

	// Load environment variables.
	errs := tpl.ParseEnvironment()
	if len(errs) > 0 {
		return gen.NewMultiError(errs...)
	}

	// Prompt user for additional information.
	answers, err := tpl.ParsePrompts()
	if err != nil {
		return err
	}

	// Consolidate the data from prompts and environment variables.
	data := make(map[string]interface{}, 0)
	for k, v := range answers {
		data[k] = v
	}
	for k, v := range tpl.Environment {
		data[k] = v
	}
	// Override.
	data["Pkg"] = pkgName
	data["PKG"] = pkgName

	isGoFile := func(name string) bool {
		return path.Ext(name) == ".go"
	}

	clone := func(act *gen.Action) error {
		src, dst := act.Source(), act.Destination()
		r, err := g.ReadOnlyFile(src)
		if err != nil {
			return err
		}

		defer func() {
			if err := r.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		if err := g.Touch(dst); err != nil {
			return err
		}

		w, err := g.WriteOnlyFile(dst)
		if err != nil {
			g.Remove(dst)
			return err
		}

		defer func() {
			if dryRun {
				return
			}
			if err := w.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		parser := func(b []byte) ([]byte, error) {
			if len(b) == 0 {
				return nil, fmt.Errorf("%s: file empty", src)
			}
			// Required.
			act.Variables["PKG"] = pkgName
			act.Variables["Pkg"] = pkgName
			b, err = gen.ParseTemplate(b, act.Variables)
			if err != nil {
				return nil, err
			}

			if isGoFile(dst) {
				b, err = gen.FormatSource(b)
				if err != nil {
					_ = g.Remove(dst)
					return nil, err
				}
			}
			return b, nil
		}

		return g.Copy(r, w, parser)
	}

	merr := gen.NewMultiError()
	for _, act := range tpl.Actions {
		newAct, err := act.Resolve(data)
		if merr.Add(err) {
			continue
		}
		if added := merr.Add(clone(newAct)); !added {
			fmt.Printf("%s: file created\n", newAct.Destination())
		}
	}
	if merr.HasError() {
		return merr
	}
	return nil
}
