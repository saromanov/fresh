package pkg

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/google/go-github/v28/github"
)

type NewRelease struct {
	Name string
	Tag  string
}

// NewReleases retruns list of all new releases
func NewReleases(deps []Dependency) ([]*NewRelease, error) {
	nr := []*NewRelease{}
	client := github.NewClient(nil)
	for _, dep := range deps {
		releases, _, err := client.Repositories.ListReleases(context.TODO(), dep.Author, dep.Name, nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get release: %v", err)
		}

		for _, r := range releases {
			v1, _ := semver.NewVersion(r.GetTagName())
			v2, _ := semver.NewVersion(dep.Tag)
			if v1.Compare(v2) == 1 {
				nr = append(nr, &NewRelease{
					Tag: r.GetTagName(),
					Name: dep.Name,
				})
				break
			}
		}
	}

	return nr, nil
}
