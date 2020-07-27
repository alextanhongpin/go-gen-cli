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
		g := gen.New(c.String("file"))
		g.SetDryRun(c.Bool("dry-run"))

		cfg, err := g.LoadConfig()
		if err != nil {
			return err
		}

		pkg := c.Args().First()
		name := c.String("template")
		if name == "" {
			return errors.New("name is required: gen generate <name>")
		}
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

		answers["Pkg"] = pkg

		for key, val := range tpl.Environment {
			val, err := gen.ParseString(val, answers)
			if err != nil {
				return err
			}
			answers[key] = val
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
				g.Remove(out)
				return err
			}

			defer func() {
				if c.Bool("dry-run") {
					return
				}
				if err := w.Close(); err != nil {
					log.Fatal(err)
				}
			}()

			parser := func(b []byte) ([]byte, error) {
				if len(b) == 0 {
					return nil, fmt.Errorf("%s: file empty", in)
				}
				b, err = gen.ParseTemplate(b, answers)
				if err != nil {
					return nil, err
				}

				if isGoFile(out) {
					b, err = gen.FormatSource(b)
					if err != nil {
						_ = g.Remove(out)
						return nil, err
					}
				}
				return b, nil
			}

			return g.Copy(r, w, parser)
		}

		merr := gen.NewMultiError()
		for _, vol := range tpl.Volumes {
			src, dst, err := vol.Split()
			if err != nil {
				return err
			}
			dst, err = gen.ParseString(dst, answers)
			if merr.Add(err) {
				continue
			}
			if added := merr.Add(copyTo(src, dst)); !added {
				fmt.Printf("%s: file created\n", dst)
			}
		}
		if merr.HasError() {
			return merr
		}
		return nil
	},
}
