package pkg

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Update provides updating of deps.
// It might be updated specific dep or all deps
func Update(path string, releases []*NewRelease) error {
	if path == "" {
		return fmt.Errorf("path is not defined")
	}
	url := func(r *NewRelease) string {
		res := strings.Split(r.URL, "https://")[1]
		return fmt.Sprintf("%s@%s", res, r.Tag)
	}
	for _, r := range releases {
		if err := update(url(r)); err != nil {
			fmt.Printf("unable to update release: %v", err)
			continue
		}
		Infof("release updated: %s %s", r.Name, r.Tag)
	}
	return nil
}

func update(path string) error {
	if path == "" {
		return fmt.Errorf("path is not defined")
	}
	cmd := exec.Command("go", "get", "-u", path)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to exec command: %v", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	return nil
}
