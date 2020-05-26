package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

type Data struct {
	Prompt map[string]interface{}
	Env    map[string]string
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
		log.Fatal(err)
		os.Exit(1)
	}
}
