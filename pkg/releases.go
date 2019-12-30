package internal

import (
	"context"
	"fmt"

	"github.com/google/go-github/v28/github"
)

func NewReleases(deps []Dependency) error {
	client := github.NewClient(nil)
	for _, dep := range deps {
		releases, _, err := client.Repositories.ListReleases(context.TODO(), dep.Name, dep.Name, nil)
		if err != nil {
			return fmt.Errorf("unable to get release: %v", err)
		}

		for _, r := range releases {
			fmt.Println(r.GetTagName())
		}
	}

	return nil
}
