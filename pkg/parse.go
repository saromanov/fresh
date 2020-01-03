package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

// Dependency defines single dependency
// in format name-tag
type Dependency struct {
	Path   string
	Tag    string
	Author string
	Name   string
}

// Parse provides parsing of go.mod file
func Parse(path string) ([]Dependency, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []Dependency{}, fmt.Errorf("unable to open file: %v", err)
	}

	return parse(data)
}

func parse(data []byte) ([]Dependency, error) {
	str := string(data)
	lines := strings.Split(str, "\n")
	response := []Dependency{}
	for _, line := range lines {
		// TODO: Need to generalize
		trim := strings.TrimSpace(line)
		if strings.HasPrefix(trim, "github") {
			item := strings.Split(trim, " ")
			author, name, err := getRepoAuthorAndName(item[0])
			if err != nil {
				return []Dependency{}, err
			}
			response = append(response, Dependency{
				Path:   item[0],
				Tag:    item[1],
				Author: author,
				Name:   name,
			})
		}
	}

	return response, nil
}

// retruns author of the repo from url and name
func getRepoAuthorAndName(s string) (string, string, error) {
	result := strings.Split(s, "/")
	if len(result) < 2 {
		return "", "", errors.New("unable to get author of the repo")
	}
	return result[1], result[2], nil
}
