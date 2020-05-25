package main

import (
	"os"
	"sort"

	"github.com/alextanhongpin/go-gen/pkg/gen"
	"github.com/urfave/cli"
)

type Data struct {
	Prompt map[string]interface{}
	Env    map[string]interface{}
}

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Value:   "gen.yaml",
			Usage:   "Load configuration from file",
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
			lsCmd,
			clearCmd,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		gen.Error(err.Error())
		os.Exit(1)
	}
}
