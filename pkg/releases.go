package pkg

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

// NewRelease defines struct for showing new releases
type NewRelease struct {
	Name        string
	Tag         string
	CurrentTag  string
	PublishedAt string
	Body        string
	URL         string
}

// NewReleases retruns list of all new releases
func NewReleases(token string, deps []Dependency) ([]*NewRelease, error) {
	nr := []*NewRelease{}
	client := getGithubClient(token)
	for _, dep := range deps {
		releases, _, err := client.Repositories.ListReleases(context.TODO(), dep.Author, dep.Name, nil)
		if err != nil {
			return nil, fmt.Errorf("unable to get release: %v", err)
		}

		for _, r := range releases {
			tag := r.GetTagName()
			equal, err := compareReleases(tag, dep.Tag)
			if err != nil {
				return nil, fmt.Errorf("unable to compare releases: %v", err)
			}
			if equal {
				nr = append(nr, &NewRelease{
					Tag:         tag,
					Name:        dep.Name,
					CurrentTag:  dep.Tag,
					PublishedAt: r.GetPublishedAt().String(),
					Body:        r.GetBody(),
					URL:         prepareURL(r.GetHTMLURL()),
				})
				break
			}
		}
	}

	return nr, nil
}

func compareReleases(tag, depTag string) (bool, error) {
	v1, err := semver.NewVersion(tag)
	if err != nil {
		return false, fmt.Errorf("unable to validate version: %v", err)
	}
	v2, err := semver.NewVersion(depTag)
	if err != nil {
		return false, fmt.Errorf("unable to validate version: %v", err)
	}
	return v1.Compare(v2) == 1, nil
}

// returns Github client
// if env.variables contains GITHUB_TOKEN
// it retruns authentificated client
// or it contains from command line variables
func getGithubClient(tokenInp string) *github.Client {
	token := getGithubToken(tokenInp)
	if token == "" {
		return github.NewClient(nil)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func getGithubToken(token string) string {
	if token == "" {
		return os.Getenv("GITHUB_TOKEN")
	}
	return token
}

// prepareURL returns url to the repo
func prepareURL(path string) string {
	if !strings.Contains(path, "/releases") {
		return path
	}
	return strings.Split(path, "/releases")[0]
}
