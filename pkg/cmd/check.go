package cmd

import (
	"fmt"

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
		return nil
	}

	for _, r := range releases {
		pkg.Info(fmt.Sprintf("New version of: %s %s", r.Name, r.Tag))
	}
	return nil
}
