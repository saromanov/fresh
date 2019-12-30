package cmd

import "github.com/saromanov/fresh/pkg"

// Check provides checking of directory with go modules
func Check(path string) error {
	deps, err := pkg.Parse(path)
	if err != nil {
		panic(err)
	}

	pkg.NewReleases(deps)
	return nil
}
