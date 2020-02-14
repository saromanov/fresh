package pkg

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Update provides updating of deps.
// It might be updated specific dep or all deps
func Update(path string) error {
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
