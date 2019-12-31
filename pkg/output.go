package pkg

import (
	"github.com/fatih/color"
)

// Info prints text in blue
func Info(s string) {
	color.Blue(s)
}