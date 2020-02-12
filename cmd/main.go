package main

import (
	"os"

	"github.com/saromanov/fresh/pkg/cmd"
)

func main() {
	cmd.Check(os.Args)
}
