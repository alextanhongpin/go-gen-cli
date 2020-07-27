package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	var version string
	b, err := ioutil.ReadFile("VERSION")
	if err != nil {
		version = "unknown"
	}
	version = string(b)

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
		Version: version,
		Flags:   flags,
		Commands: []*cli.Command{
			initCmd,
			generateCmd,
			clearCmd,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
