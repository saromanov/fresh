package main

import (
	"os"

	"github.com/saromanov/fresh/pkg/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "fresh",
		Usage: "Checking of newest deps",
		Commands: []*cli.Command{
			{
				Name:    "check",
				Aliases: []string{"c"},
				Usage:   "starting of checking",
				Action: func(c *cli.Context) error {
					cmd.Check("go.mod")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
