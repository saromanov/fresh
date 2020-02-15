package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/saromanov/fresh/pkg"
	"github.com/urfave/cli/v2"
)

// Check is entry point for the app
func Check(args []string) {
	app := &cli.App{
		Name:  "fresh",
		Usage: "Checking of newest deps",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "update-all",
				Usage: "updating all depencencies",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "check",
				Usage: "starting of checking",
				Action: func(c *cli.Context) error {
					if err := check(c, "go.mod"); err != nil {
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
func check(c *cli.Context, path string) error {
	releases, err := getReleases(path)
	if err != nil {
		return err
	}
	if releases == nil || len(releases) == 0 {
		log.Printf("all dependencies is up to date")
		return nil
	}

	for _, r := range releases {
		pkg.Warningf("'%s'", r.Name)
		pkg.Infof("url: %s", r.URL)
		pkg.Infof("current version: %s", pkg.Text(r.CurrentTag))
		pkg.Infof("new version %s", pkg.Text(r.Tag))
		pkg.Infof("published at %s", pkg.Text(r.PublishedAt))
		pkg.Infof("release description:\n %s", pkg.Text(r.Body))
		fmt.Println()
	}
	if c.Bool("update-all") {
		pkg.Info("updating of dependencies...")
		pkg.Update(path)
	}
	return nil
}

func update(path string) error {
	releases, err := getReleases(path)
	if err != nil {
		return err
	}
	if releases == nil || len(releases) == 0 {
		log.Printf("all dependencies is up to date")
		return nil
	}

	return pkg.Update(path)
}

func getReleases(path string) ([]*pkg.NewRelease, error) {
	deps, err := pkg.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("unable to parse go.mod file: %v", err)
	}

	releases, err := pkg.NewReleases(deps)
	if err != nil {
		return nil, fmt.Errorf("unable to get new releases: %v", err)
	}
	return releases, nil
}
