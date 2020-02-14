package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/saromanov/fresh/pkg"
	"github.com/urfave/cli/v2"
)


https://medium.com/faun/managing-dependency-and-module-versioning-using-go-modules-c7c6da00787a
// CHeck is entry point for the app
func Check(args []string) {
	app := &cli.App{
		Name:  "fresh",
		Usage: "Checking of newest deps",
		Commands: []*cli.Command{
			{
				Name:  "check",
				Usage: "starting of checking",
				Action: func(c *cli.Context) error {
					if err := check("go.mod"); err != nil {
						log.Fatalf("unable to check: %v", err)
					}
					return nil
				},
			},
			{
				Name:  "update",
				Usage: "updating of deps",
				Action: func(c *cli.Context) error {
					if err := update("go.mod"); err != nil {
						log.Fatalf("unable to check: %v", err)
					}
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

// check provides checking of directory with go modules
func check(path string) error {
	deps, err := pkg.Parse(path)
	if err != nil {
		return fmt.Errorf("unable to parse go.mod file: %v", err)
	}

	releases, err := pkg.NewReleases(deps)
	if err != nil {
		return fmt.Errorf("unable to get new releases: %v", err)
	}

	if releases == nil || len(releases) == 0 {
		log.Printf("all dependencies is up to date")
		return nil
	}

	for _, r := range releases {
		pkg.Warningf("'%s'", r.Name)
		pkg.Infof("current version: %s", pkg.Text(r.CurrentTag))
		pkg.Infof("new version %s", pkg.Text(r.Tag))
		pkg.Infof("published at %s", pkg.Text(r.PublishedAt))
		pkg.Infof("release description:\n %s", pkg.Text(r.Body))
		fmt.Println()
	}
	return nil
}

func update(path string) error {
	return nil
}
