package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/alextanhongpin/go-gen"

	"github.com/urfave/cli"
)

var generateCmd = &cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Usage:   "generates the given template",
	Action: func(c *cli.Context) error {
		g := gen.New(c.String("file"))
		cfg, err := g.LoadConfig()
		if err != nil {
			return err
		}

		name := c.Args().First()
		tpl := cfg.Find(name)
		if tpl == nil {
			return fmt.Errorf("%s: not found", name)
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

		data := Data{
			Prompt: answers,
			Env:    tpl.Environment,
		}

		isGoFile := func(name string) bool {
			return path.Ext(name) == ".go"
		}

		copyTo := func(in, out string) error {
			r, err := g.ReadOnlyFile(in)
			if err != nil {
				return err
			}
			defer func() {
				if err := r.Close(); err != nil {
					log.Fatal(err)
				}
			}()

			if err := g.Touch(out); err != nil {
				return err
			}
			w, err := g.WriteOnlyFile(out)
			if err != nil {
				return err
			}
			defer func() {
				if err := w.Close(); err != nil {
					log.Fatal(err)
				}
			}()

			parser := func(b []byte) ([]byte, error) {
				if len(b) == 0 {
					return nil, fmt.Errorf("%s: file empty", in)
				}
				b, err = gen.ParseTemplate(b, data)
				if err != nil {
					return nil, err
				}

				if isGoFile(out) {
					b, err = gen.FormatSource(b)
					if err != nil {
						return nil, err
					}
				}
				return b, nil
			}

			return g.Copy(r, w, parser)
		}

		merr := gen.NewMultiError()
		for _, act := range tpl.Actions {
			in := os.ExpandEnv(act.Template)
			out := os.ExpandEnv(act.Path)
			if added := merr.Add(copyTo(in, out)); !added {
				fmt.Printf("%s: file created\n", out)
			}
		}
		return merr
	},
}
