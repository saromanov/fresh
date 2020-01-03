package cmd

import (
	"fmt"
	"log"

	"github.com/saromanov/fresh/pkg"
)

// Check provides checking of directory with go modules
func Check(path string) error {
	deps, err := pkg.Parse(path)
	if err != nil {
		return fmt.Errorf("unable to parse go.mod file: %v", err)
	}

	releases, err := pkg.NewReleases(deps)
	if err != nil {
		return fmt.Errorf("unable to get new releases: %v", err)
	}

	if releases == nil || len(releases) == 0 {
		log.Infof("all dependencies is up to date")
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
