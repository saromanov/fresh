package pkg

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Update provides updating of deps.
// It might be updated specific dep or all deps
func Update(path string, releases []*NewRelease) error {
	if path == "" {
		return fmt.Errorf("path is not defined")
	}
	url := func(r *NewRelease) string {
		return fmt.Sprintf("%s@%s", r.URL, r.Tag)
	}
	for _, r := range releases {
		if err := update(url(r)); err != nil {
			return fmt.Errorf("unable to update release: %v", err)
		}
	}
	return nil
}

func update(path string) error {
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
