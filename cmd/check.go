package check

import (
	"github.com/saromanov/fresh/internal"
	"github.com/urfave/cli/v2"
)

func Check(c *cli.Context) error {
	deps, err := internal.Parse("go.mod")
	if err != nil {
		panic(err)
	}

	internal.NewReleases(deps)
	return nil
}
