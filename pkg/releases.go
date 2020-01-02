package pkg

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/google/go-github/v28/github"
)

// NewRelease defines struct for showing new releases
type NewRelease struct {
	Name        string
	Tag         string
	CurrentTag  string
	PublishedAt string
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
			tag := r.GetTagName()
			v1, err := semver.NewVersion(tag)
			if err != nil {
				return nil, fmt.Errorf("unable to validate version: %v", err)
			}
			v2, err := semver.NewVersion(dep.Tag)
			if err != nil {
				return nil, fmt.Errorf("unable to validate version: %v", err)
			}
			if v1.Compare(v2) == 1 {
				nr = append(nr, &NewRelease{
					Tag:         tag,
					Name:        dep.Name,
					CurrentTag:  dep.Tag,
					PublishedAt: r.GetPublishedAt().String(),
				})
				break
			}
		}
	}

	return nr, nil
}
