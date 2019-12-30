package internal

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Dependency defines single dependency
// in format name-tag
type Dependency struct {
	Name string
	Tag  string
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
		if strings.HasPrefix(line, "git") {
			item := strings.Split(line, " ")
			response = append(response, Dependency{
				Name: item[0],
				Tag:  item[1],
			})
		}
	}

	return response, nil
}
