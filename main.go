package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/tj/survey"
	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"
	"golang.org/x/tools/imports"
	"gopkg.in/yaml.v2"
)

type data struct {
	Type string
	Name string
}

type Config struct {
	Commands map[string]*Command `yaml:"commands"`
}

type Command struct {
	Description string `yaml:"description"`
	Template    string `yaml:"template"`
	Path        string `yaml:"path"`
}

// the questions to ask
var qs = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "What is your name?"},
		Validate: survey.Required,
	},
}

func open(fname string, flag int) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dirpath := filepath.Join(dir, path.Dir(fname))
	if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	filepath := filepath.Join(dir, fname)
	file, err := os.OpenFile(filepath, flag, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

var app = cli.NewApp()

func main() {
	app.Name = "Gen CLI"
	app.Usage = "CLI for code generation"
	app.Authors = []*cli.Author{
		{Name: "Alex Tan", Email: "alextan220990@gmail.com"},
	}
	app.Version = "1.0.0"
	flags := []cli.Flag{
		altsrc.NewIntFlag(&cli.IntFlag{Name: "test"}),
		&cli.StringFlag{Name: "load"},
	}
	app.Commands = []*cli.Command{
		{

			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init a config gen.yml",
			Action: func(c *cli.Context) error {
				w, err := open("gen.yml", os.O_WRONLY|os.O_CREATE|os.O_EXCL)
				if err != nil {
					return err
				}
				defer w.Close()

				t := make(map[string]map[string]interface{})
				t["commands"] = make(map[string]interface{})
				t["commands"]["test"] = &Command{
					Description: "Describe this action",
					Template:    "Path to template",
					Path:        "Path to destination",
				}
				d, err := yaml.Marshal(&t)
				if err != nil {
					return err
				}
				log.Println(string(d))
				n, err := w.Write(d)
				if err != nil {
					return err
				}
				fmt.Printf("Write %d bytes", n)
				return nil
			},
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate a reusable template",
			Action: func(c *cli.Context) error {
				// Load the config.yaml.
				rw, err := open("gen.yml", os.O_RDWR)
				if err != nil {
					return err
				}
				defer rw.Close()

				data, err := ioutil.ReadAll(rw)
				if err != nil {
					return err
				}

				t := make(map[string]map[string]*Command)
				if err := yaml.Unmarshal(data, &t); err != nil {
					return err
				}
				log.Println("read from yaml", t)

				// Check if the command exists, if yes, skip.
				// Else, create the command with a basic template.
				command := c.Args().First()
				_, ok := t["commands"][command]
				if !ok {
					t["commands"][command] = &Command{
						Description: fmt.Sprintf("Generates a %s", command),
						Template:    fmt.Sprintf("templates/%s.go", command),
						Path:        "Enter your path here",
					}
					d, err := yaml.Marshal(&t)
					if err != nil {
						return err
					}
					n, err := rw.WriteAt(d, 0)
					if err != nil {
						return err
					}
					fmt.Printf("Write %d bytes", n)
					return nil
				}
				return nil
				// fmt.Println("added task: ", c.Args().First())
				// answers := make(map[string]interface{})
				// // perform the questions
				// if err := survey.Ask(qs, &answers); err != nil {
				//         return err
				// }
				// fmt.Println("executing generate", answers)
				// return nil
			},
			Before: altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("load")),
			Flags:  flags,
		},
		{
			Name:    "exec",
			Aliases: []string{"e"},
			Usage:   "executes a template",
			Action: func(c *cli.Context) error {
				r, err := open("gen.yml", os.O_RDONLY)
				if err != nil {
					return err
				}
				defer r.Close()

				data, err := ioutil.ReadAll(r)
				if err != nil {
					return err
				}

				var cfg Config
				if err := yaml.Unmarshal(data, &cfg); err != nil {
					return err
				}

				command := c.Args().First()
				cmd, ok := cfg.Commands[command]
				if !ok {
					return errors.New("command not found")
				}
				// Pass in pkg and name.
				return core(cmd.Template, cmd.Path)
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// FormatSource is gofmt with addition of removing any unused imports.
func FormatSource(source []byte) ([]byte, error) {
	return imports.Process("", source, &imports.Options{
		AllErrors: true, Comments: true, TabIndent: true, TabWidth: 8,
	})
}

func core(in, out string) error {
	// Open as read-only, create if not exists.
	r, err := open(in, os.O_RDONLY|os.O_CREATE|os.O_EXCL)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			r, err = os.Open(in)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	defer r.Close()

	tpl, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if len(string(tpl)) == 0 {
		return errors.New("template is empty")
	}
	t := template.Must(template.New("").Parse(string(tpl)))

	// Open as write-only, create if not exists.
	w, err := open(out, os.O_WRONLY|os.O_CREATE|os.O_EXCL)
	if err != nil {
		return err
	}
	defer w.Close()

	var d map[string]interface{}
	var b bytes.Buffer
	if err := t.Execute(&b, d); err != nil {
		return err
	}
	b2, err := FormatSource(b.Bytes())
	if err != nil {
		return err
	}
	if _, err := w.Write(b2); err != nil {
		return err
	}
	return nil
}
