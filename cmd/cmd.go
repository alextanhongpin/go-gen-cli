package main

import (
	"log"
	"os"
	"sort"

	"github.com/alextanhongpin/go-gen/pkg/gen"

	"github.com/urfave/cli"
)

type Data struct {
	PackageName string
	StructName  string
	Tag         string
	Template    *gen.Template
	Prompts     map[string]interface{}
}

var data Data

var cfgPath string

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Value:       "gen.yaml",
			Usage:       "Load configuration from file",
			Destination: &cfgPath,
		},
	}

	app := &cli.App{
		Name:  "Gen CLI",
		Usage: "CLI for code generation",
		Authors: []*cli.Author{
			{Name: "Alex Tan", Email: "alextan220990@gmail.com"},
		},
		Version: "1.0.0",
		Flags:   flags,
		Commands: []*cli.Command{
			initCmd,
			addCmd,
			removeCmd,
			generateCmd,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
