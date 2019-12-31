package pkg

import (
	"context"
	"fmt"

	"github.com/blang/semver"
	"github.com/google/go-github/v28/github"
)

type NewRelease struct {
	Name string
	Tag  string
}

// NewReleases retruns list of all new releases
func NewReleases(deps []Dependency) ([]*NewRelease, error) {
	client := github.NewClient(nil)
	for _, dep := range deps {
		releases, _, err := client.Repositories.ListReleases(context.TODO(), dep.Author, dep.Name, nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get release: %v", err)
		}

		for _, r := range releases {
			v1, _ := semver.Make(r.GetTagName())
			v2, _ := semver.Make(dep.Tag)
			fmt.Println(v1.Equals(v2), r.GetTagName(), dep.Tag)
		}
	}

	return nil, nil
}
